package samples

import "recsys/model"

func Sample5() model.Problem {
	return model.Problem{
		Clubs: []model.Club{
			{ID: 1}, {ID: 2}, {ID: 3},
			{ID: 4}, {ID: 5}, {ID: 6},
		},
		Events: []model.Event{
			// Club 1 Events
			model.NewEvent(1, 1, 0, 3),
			model.NewEvent(2, 1, 4, 6),
			model.NewEvent(3, 1, 7, 9),
			model.NewEvent(4, 1, 10, 12),

			// Club 2 Events
			model.NewEvent(5, 2, 1, 3),
			model.NewEvent(6, 2, 4, 7),
			model.NewEvent(7, 2, 8, 10),
			model.NewEvent(8, 2, 11, 13),

			// Club 3 Events
			model.NewEvent(9, 3, 0, 2),
			model.NewEvent(10, 3, 3, 6),
			model.NewEvent(11, 3, 7, 10),
			model.NewEvent(12, 3, 11, 14),

			// Club 4 Events
			model.NewEvent(13, 4, 2, 5),
			model.NewEvent(14, 4, 6, 9),
			model.NewEvent(15, 4, 10, 13),
			model.NewEvent(16, 4, 14, 17),

			// Club 5 Events
			model.NewEvent(17, 5, 1, 4),
			model.NewEvent(18, 5, 5, 7),
			model.NewEvent(19, 5, 8, 11),
			model.NewEvent(20, 5, 12, 15),

			// Club 6 Events
			model.NewEvent(21, 6, 0, 3),
			model.NewEvent(22, 6, 4, 6),
			model.NewEvent(23, 6, 7, 10),
			model.NewEvent(24, 6, 11, 14),

			// Mixed overlapping events for testing
			model.NewEvent(25, 1, 15, 17),
			model.NewEvent(26, 2, 16, 18),
			model.NewEvent(27, 3, 17, 19),
			model.NewEvent(28, 4, 18, 20),
			model.NewEvent(29, 5, 19, 20),
			model.NewEvent(30, 6, 15, 20),
		},
	}
}
