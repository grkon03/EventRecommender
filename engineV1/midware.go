package engineV1

import (
	"recsys/model"
	"recsys/util"
)

func devideBySectionStart(p model.Problem) map[int]util.Set[model.Event] {
	sections := make(map[int]util.Set[model.Event])

	for _, e := range p.Events {
		s := e.SectionStart
		if _, exist := sections[s]; !exist {
			sections[s] = util.NewSet[model.Event]()
		}
		sections[s].Add(e)
	}

	return sections
}
