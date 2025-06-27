package samples

import "recsys/model"

// no longer able to calculate by full search...
func Sample2() model.Problem {
	return model.Problem{
		Clubs: []model.Club{
			{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5}, {ID: 6}, {ID: 7},
		},
		Events: []model.Event{
			// Club 1 Events
			model.NewEvent(1, 1, 0, 2),
			model.NewEvent(2, 1, 3, 6),
			model.NewEvent(3, 1, 7, 9),
			model.NewEvent(4, 1, 10, 14),

			// Club 2 Events
			model.NewEvent(5, 2, 0, 1),
			model.NewEvent(6, 2, 2, 5),
			model.NewEvent(7, 2, 6, 8),
			model.NewEvent(8, 2, 9, 12),

			// Club 3 Events
			model.NewEvent(9, 3, 1, 3),
			model.NewEvent(10, 3, 4, 7),
			model.NewEvent(11, 3, 8, 10),
			model.NewEvent(12, 3, 11, 15),

			// Club 4 Events
			model.NewEvent(13, 4, 0, 3),
			model.NewEvent(14, 4, 4, 8),
			model.NewEvent(15, 4, 9, 13),
			model.NewEvent(16, 4, 14, 17),

			// Club 5 Events
			model.NewEvent(17, 5, 1, 4),
			model.NewEvent(18, 5, 5, 9),
			model.NewEvent(19, 5, 10, 13),
			model.NewEvent(20, 5, 14, 16),

			// Club 6 Events
			model.NewEvent(21, 6, 0, 2),
			model.NewEvent(22, 6, 3, 5),
			model.NewEvent(23, 6, 6, 9),
			model.NewEvent(24, 6, 10, 14),

			// Club 7 Events
			model.NewEvent(25, 7, 1, 3),
			model.NewEvent(26, 7, 4, 6),
			model.NewEvent(27, 7, 7, 10),
			model.NewEvent(28, 7, 11, 15),
		},
	}
}
