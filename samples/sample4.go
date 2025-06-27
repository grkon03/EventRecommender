package samples

import "recsys/model"

func Sample4() model.Problem {
	return model.Problem{
		Clubs: []model.Club{
			{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4},
			{ID: 5}, {ID: 6}, {ID: 7}, {ID: 8},
		},
		Events: []model.Event{
			// Club 1 Events
			model.NewEvent(1, 1, 0, 3),
			model.NewEvent(2, 1, 4, 6),
			model.NewEvent(3, 1, 7, 10),
			model.NewEvent(4, 1, 11, 13),
			model.NewEvent(5, 1, 14, 16),
			model.NewEvent(6, 1, 18, 21),

			// Club 2 Events
			model.NewEvent(7, 2, 1, 4),
			model.NewEvent(8, 2, 5, 8),
			model.NewEvent(9, 2, 9, 11),
			model.NewEvent(10, 2, 12, 15),
			model.NewEvent(11, 2, 16, 19),
			model.NewEvent(12, 2, 20, 23),

			// Club 3 Events
			model.NewEvent(13, 3, 2, 5),
			model.NewEvent(14, 3, 6, 9),
			model.NewEvent(15, 3, 10, 12),
			model.NewEvent(16, 3, 13, 16),
			model.NewEvent(17, 3, 17, 20),
			model.NewEvent(18, 3, 21, 24),

			// Club 4 Events
			model.NewEvent(19, 4, 0, 2),
			model.NewEvent(20, 4, 3, 6),
			model.NewEvent(21, 4, 7, 9),
			model.NewEvent(22, 4, 10, 13),
			model.NewEvent(23, 4, 14, 17),
			model.NewEvent(24, 4, 18, 20),

			// Club 5 Events
			model.NewEvent(25, 5, 2, 4),
			model.NewEvent(26, 5, 5, 7),
			model.NewEvent(27, 5, 8, 11),
			model.NewEvent(28, 5, 12, 14),
			model.NewEvent(29, 5, 15, 18),
			model.NewEvent(30, 5, 19, 22),

			// Club 6 Events
			model.NewEvent(31, 6, 1, 3),
			model.NewEvent(32, 6, 4, 6),
			model.NewEvent(33, 6, 7, 10),
			model.NewEvent(34, 6, 11, 13),
			model.NewEvent(35, 6, 14, 17),
			model.NewEvent(36, 6, 18, 21),

			// Club 7 Events
			model.NewEvent(37, 7, 0, 3),
			model.NewEvent(38, 7, 4, 7),
			model.NewEvent(39, 7, 8, 10),
			model.NewEvent(40, 7, 11, 14),
			model.NewEvent(41, 7, 15, 18),
			model.NewEvent(42, 7, 19, 22),

			// Club 8 Events
			model.NewEvent(43, 8, 1, 4),
			model.NewEvent(44, 8, 5, 8),
			model.NewEvent(45, 8, 9, 12),
			model.NewEvent(46, 8, 13, 16),
			model.NewEvent(47, 8, 17, 20),
			model.NewEvent(48, 8, 21, 24),

			// Overlapping edge cases for testing
			model.NewEvent(49, 4, 22, 25),
			model.NewEvent(50, 8, 22, 25),
		},
	}
}
