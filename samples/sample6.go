package samples

import "recsys/model"

func Sample6() model.Problem {
	return model.Problem{
		Clubs: []model.Club{
			{ID: 1}, {ID: 2}, {ID: 3},
			{ID: 4}, {ID: 5}, {ID: 6},
			{ID: 7}, {ID: 8},
		},
		Events: []model.Event{
			// Club 1: Long and overlapping events
			model.NewEvent(1, 1, 0, 10), // Long event
			model.NewEvent(2, 1, 5, 12), // Overlaps with Event 1
			model.NewEvent(3, 1, 15, 20),
			model.NewEvent(4, 1, 18, 25), // Overlaps with Event 3

			// Club 2: Short and evenly spaced events
			model.NewEvent(5, 2, 0, 2),
			model.NewEvent(6, 2, 3, 5),
			model.NewEvent(7, 2, 6, 8),
			model.NewEvent(8, 2, 9, 11),
			model.NewEvent(9, 2, 12, 15),
			model.NewEvent(10, 2, 16, 18),

			// Club 3: Mixed lengths
			model.NewEvent(11, 3, 1, 4),
			model.NewEvent(12, 3, 5, 9), // Long event
			model.NewEvent(13, 3, 10, 13),
			model.NewEvent(14, 3, 14, 17),
			model.NewEvent(15, 3, 20, 30), // Very long event

			// Club 4: Clustered and overlapping
			model.NewEvent(16, 4, 0, 3),
			model.NewEvent(17, 4, 2, 6), // Overlaps with Event 16
			model.NewEvent(18, 4, 5, 9), // Overlaps with Event 17
			model.NewEvent(19, 4, 10, 15),
			model.NewEvent(20, 4, 14, 20), // Overlaps with Event 19

			// Club 5: Sparse but long events
			model.NewEvent(21, 5, 0, 8),   // Long event
			model.NewEvent(22, 5, 10, 20), // Long event
			model.NewEvent(23, 5, 21, 30), // Long event

			// Club 6: Alternating lengths
			model.NewEvent(24, 6, 1, 3),
			model.NewEvent(25, 6, 5, 7),
			model.NewEvent(26, 6, 8, 15), // Long event
			model.NewEvent(27, 6, 16, 18),
			model.NewEvent(28, 6, 19, 25),

			// Club 7: Dense and overlapping
			model.NewEvent(29, 7, 0, 4),
			model.NewEvent(30, 7, 3, 7),  // Overlaps with Event 29
			model.NewEvent(31, 7, 6, 10), // Overlaps with Event 30
			model.NewEvent(32, 7, 12, 14),
			model.NewEvent(33, 7, 15, 17),
			model.NewEvent(34, 7, 18, 22),

			// Club 8: Isolated and long events
			model.NewEvent(35, 8, 0, 5),
			model.NewEvent(36, 8, 6, 11),
			model.NewEvent(37, 8, 12, 17),
			model.NewEvent(38, 8, 18, 30), // Very long event

			// Mixed rare cases
			model.NewEvent(39, 1, 25, 30), // Long tail overlap
			model.NewEvent(40, 2, 19, 30), // Long tail overlap
			model.NewEvent(41, 3, 0, 1),   // Very short
			model.NewEvent(42, 4, 19, 22), // Isolated cluster
			model.NewEvent(43, 5, 3, 6),   // Small gap filler
			model.NewEvent(44, 6, 14, 16), // Small overlap
			model.NewEvent(45, 7, 20, 22), // Gap filler
			model.NewEvent(46, 8, 7, 9),   // Small event
			model.NewEvent(47, 1, 13, 14), // Single gap filler
			model.NewEvent(48, 2, 8, 9),   // Short inside range
			model.NewEvent(49, 3, 18, 19), // Small gap
			model.NewEvent(50, 4, 22, 23), // Tiny overlap
		},
	}
}
