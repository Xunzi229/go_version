package go_version

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type GoVersion struct {
	Version string
}
var (
	VersionPattern         = `[0-9]+(?>\.[0-9a-zA-Z]+)*(-[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*)?` // :nodoc:
  	AnchoredVersionPattern = `\A\s*(`+VersionPattern+`)?\s*\z`               // :nodoc:
)

func New(version string) GoVersion {
	correct(version)

	v := version
	if len(version) == 0 {
		v = "0"
	}
	reg := regexp.MustCompile("-")

	v = strings.TrimSpace(v)
	v = reg.ReplaceAllString(v, ".pre.")

	return GoVersion{
		Version: v,
	}
}

func (g GoVersion)Compare(v string, compare string)  {

}

// 该版本大于等于v1
func (g GoVersion)ge(v1 string){
}

// 该版本大于v1
func (g GoVersion)gt(v1 string){
}

// <=>
func (g GoVersion)EachEqual(other GoVersion) int{
	lhSegments := g.segments()
	rhSegments := other.segments()

	lhSize := len(lhSegments)
    rhSize := len(rhSegments)

    limit := rhSize
    if lhSize > rhSize {
    	limit = lhSize
	}
	limit--

    i := 0

    r := regexp.MustCompile(`^\d+$`)
	a := regexp.MustCompile(`[a-z]+`)
    for i < limit {
		lhs, rhs := 0, 0
		if i <= len(lhSegments) {
			if r1 := r.Match([]byte(lhSegments[i])); r1 {
				vi, _ := strconv.Atoi(lhSegments[i])
				rhs = vi
			}
		}
		if i <= len(rhSegments) {
			if r1 := r.Match([]byte(rhSegments[i])); r1 {
				vi, _ := strconv.Atoi(rhSegments[i])
				lhs = vi
			}
		}
		i++

		if lhs == rhs {
			continue
		}
		if a1 := a.Match([]byte(lhSegments[i])); a1 {
			if r1 := r.Match([]byte(rhSegments[i])); r1 {
				return -1
			}
		}

		if a1 := a.Match([]byte(rhSegments[i])); a1 {
			if r1 := r.Match([]byte(lhSegments[i])); r1 {
				return 1
			}
		}

		// 都是string
		if a1 := a.Match([]byte(rhSegments[i])); a1 {
			if a2 := a.Match([]byte(lhSegments[i])); a2 {
				return strings.Compare(rhSegments[i], lhSegments[i])
			}
		}
		if lhs > rhs {
			return 1
		} else {
			return -1
		}
	}

	return 0
}

func (g GoVersion)segments()[]string {
	gs := make([]string, 0)
	reg := regexp.MustCompile("(?i)[0-9]+|[a-z]+")
	r := regexp.MustCompile(`^\d+$`)
	for _, v := range reg.FindStringSubmatch(g.Version) {
		//if r1 := r.Match([]byte(v)); r1 {
		//	vi, _ := strconv.Atoi(v)
		//	gs = append(gs, vi)
		//	continue
		//}
		gs = append(gs, v)
	}
	return gs
}


func correct(version string) {
	reg, _ := regexp.Compile(AnchoredVersionPattern)
	if ok := reg.Match([]byte(version)); !ok {
		msg := fmt.Sprintf("Malformed version number string %s", version)
		panic(msg)
	}
}
