package dotdsl
import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)
type Properties map[string]any
type Graph struct {
	nodes map[string]Properties
	edges map[string]Properties
	attrs Properties
}
func Parse(data string) (*Graph, error) {
	data = strings.TrimSpace(data)
	if !strings.HasPrefix(data, "graph {") {
		return nil, errors.New("invalid graph")
	}
	lines := sanitize(data)
	g := &Graph{
		nodes: make(map[string]Properties),
		edges: make(map[string]Properties),
		attrs: nil,
	}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// comments
		if strings.HasPrefix(line, "//") || strings.HasPrefix(line, "#") {
			continue
		}
        if strings.Contains(line, "==") {
        	return nil, errors.New("invalid edge")
        }
		switch {
		case strings.HasPrefix(line, "["):
			if err := g.handleGraphAttr(line); err != nil {
				return nil, err
			}
		case strings.Contains(line, "--"):
			if err := g.handleEdge(line); err != nil {
				return nil, err
			}
		case strings.Contains(line, "["):
			if err := g.handleNode(line); err != nil {
				return nil, err
			}
		case strings.Contains(line, "="):
			if err := g.handleGraphAttr(line); err != nil {
				return nil, err
			}
		default:
			if err := g.handleNode(line); err != nil {
				return nil, err
			}
		}
	}
	return g, nil
}
func sanitize(data string) []string {
	data = strings.ReplaceAll(data, "graph {", "")
	data = strings.ReplaceAll(data, "}", "")
	var res []string
	for _, line := range strings.Split(data, "\n") {
		line = strings.Split(line, "//")[0]
		line = strings.Split(line, "#")[0]
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, ";")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p != "" {
				res = append(res, p)
			}
		}
	}
	return res
}
func (g *Graph) handleGraphAttr(line string) error {
	line = strings.Trim(line, "[]")
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return errors.New("invalid attribute")
	}
	k := strings.TrimSpace(parts[0])
	v := strings.TrimSpace(parts[1])
	if k == "" {
		return errors.New("invalid attribute")
	}
	if g.attrs == nil {
		g.attrs = make(Properties)
	}
	g.attrs[k] = parseValue(v)
	return nil
}
func (g *Graph) handleNode(line string) error {
	nodeSplit := strings.Split(line, "[")
	name := strings.TrimSpace(nodeSplit[0])
	if !isValidNode(name) {
		return errors.New("node name must be alphanumeric")
	}
	if _, ok := g.nodes[name]; !ok {
		g.nodes[name] = nil
	}
	if len(nodeSplit) == 1 {
		return nil
	}
	prop := strings.TrimSuffix(nodeSplit[1], "]")
	parts := strings.SplitN(prop, "=", 2)
	if len(parts) != 2 {
		return errors.New("invalid attribute")
	}
	k := strings.TrimSpace(parts[0])
	v := strings.TrimSpace(parts[1])
	if k == "" {
		return errors.New("invalid attribute")
	}
	if g.nodes[name] == nil {
		g.nodes[name] = make(Properties)
	}
	g.nodes[name][k] = parseValue(v)
	return nil
}
func (g *Graph) handleEdge(line string) error {
	parts := strings.Split(line, "[")
	raw := strings.TrimSpace(parts[0])
	nodes := strings.Split(raw, "--")
	// validate structure
	if len(nodes) < 2 {
		return errors.New("invalid edge")
	}
	// validate all segments
	for _, n := range nodes {
        n = strings.TrimSpace(n)
        if n == "" {
            return errors.New("invalid edge")
        }
        if len(strings.Fields(n)) != 1 {
            return errors.New("invalid edge")
        }
        if !isValidNode(n) {
            return errors.New("invalid edge")
        }
    }
	// build edges
	for i := 0; i < len(nodes)-1; i++ {
		a := strings.TrimSpace(nodes[i])
		b := strings.TrimSpace(nodes[i+1])
		g.ensureNode(a)
		g.ensureNode(b)
		key := edgeKey(a, b)
		if _, ok := g.edges[key]; !ok {
			g.edges[key] = nil
		}
	}
	// no attributes
	if len(parts) == 1 {
		return nil
	}
	attr := strings.TrimSuffix(parts[1], "]")
	p := strings.SplitN(attr, "=", 2)
	if len(p) != 2 {
		return errors.New("invalid attribute")
	}
	k := strings.TrimSpace(p[0])
	v := strings.TrimSpace(p[1])
	if k == "" {
		return errors.New("invalid attribute")
	}
	for i := 0; i < len(nodes)-1; i++ {
		a := strings.TrimSpace(nodes[i])
		b := strings.TrimSpace(nodes[i+1])
		key := edgeKey(a, b)
		if g.edges[key] == nil {
			g.edges[key] = make(Properties)
		}
		g.edges[key][k] = parseValue(v)
	}
	return nil
}
func (g *Graph) ensureNode(n string) {
	if _, ok := g.nodes[n]; !ok {
		g.nodes[n] = nil
	}
}
func edgeKey(a, b string) string {
	if a < b {
		return "{" + a + " " + b + "}"
	}
	return "{" + b + " " + a + "}"
}
func isValidNode(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
func parseValue(v string) any {
	v = strings.Trim(v, "\"")
	if i, err := strconv.Atoi(v); err == nil {
		return i
	}
	if v == "true" {
		return true
	}
	if v == "false" {
		return false
	}
	return v
}