package samples

import "recsys/model"

func Sample1() model.Problem {
	return model.Problem{
		Clubs: []model.Club{
			{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4},
		},
		Events: []model.Event{
			model.NewEvent(1, 1, 0, 2),
			model.NewEvent(2, 1, 5, 8),
			model.NewEvent(3, 1, 9, 10),
			model.NewEvent(4, 1, 12, 16),
			model.NewEvent(5, 2, 1, 2),
			model.NewEvent(6, 2, 3, 5),
			model.NewEvent(7, 2, 6, 8),
			model.NewEvent(8, 2, 11, 12),
			model.NewEvent(9, 3, 6, 8),
			model.NewEvent(10, 3, 9, 11),
			model.NewEvent(11, 3, 13, 14),
			model.NewEvent(12, 4, 0, 3),
			model.NewEvent(13, 4, 4, 9),
		},
	}
}
