package samples

import (
	"recsys/model"
)

// randomly
func Sample3() model.Problem {
	return model.Problem{
		Clubs: []model.Club{
			{ID: 1}, {ID: 2}, {ID: 3}, {ID: 4}, {ID: 5},
			{ID: 6}, {ID: 7}, {ID: 8}, {ID: 9}, {ID: 10},
			{ID: 11}, {ID: 12}, {ID: 13}, {ID: 14}, {ID: 15},
		},
		Events: []model.Event{
			// Club 1 Events
			model.NewEvent(1, 1, 0, 3),
			model.NewEvent(2, 1, 4, 8),
			model.NewEvent(3, 1, 9, 12),
			model.NewEvent(4, 1, 13, 17),
			model.NewEvent(5, 1, 18, 21),
			model.NewEvent(6, 1, 22, 26),
			model.NewEvent(7, 1, 27, 30),

			// Club 2 Events
			model.NewEvent(8, 2, 1, 5),
			model.NewEvent(9, 2, 6, 10),
			model.NewEvent(10, 2, 11, 15),
			model.NewEvent(11, 2, 16, 20),
			model.NewEvent(12, 2, 21, 25),
			model.NewEvent(13, 2, 26, 30),
			model.NewEvent(14, 2, 31, 35),

			// Club 3 Events
			model.NewEvent(15, 3, 2, 6),
			model.NewEvent(16, 3, 7, 11),
			model.NewEvent(17, 3, 12, 16),
			model.NewEvent(18, 3, 17, 21),
			model.NewEvent(19, 3, 22, 27),
			model.NewEvent(20, 3, 28, 32),
			model.NewEvent(21, 3, 33, 37),

			// Club 4 Events
			model.NewEvent(22, 4, 3, 7),
			model.NewEvent(23, 4, 8, 12),
			model.NewEvent(24, 4, 13, 17),
			model.NewEvent(25, 4, 18, 22),
			model.NewEvent(26, 4, 23, 28),
			model.NewEvent(27, 4, 29, 34),
			model.NewEvent(28, 4, 35, 39),

			// Club 5 Events
			model.NewEvent(29, 5, 4, 8),
			model.NewEvent(30, 5, 9, 13),
			model.NewEvent(31, 5, 14, 18),
			model.NewEvent(32, 5, 19, 23),
			model.NewEvent(33, 5, 24, 28),
			model.NewEvent(34, 5, 29, 33),
			model.NewEvent(35, 5, 34, 38),

			// Club 6 Events
			model.NewEvent(36, 6, 5, 9),
			model.NewEvent(37, 6, 10, 14),
			model.NewEvent(38, 6, 15, 19),
			model.NewEvent(39, 6, 20, 24),
			model.NewEvent(40, 6, 25, 29),
			model.NewEvent(41, 6, 30, 34),
			model.NewEvent(42, 6, 35, 39),

			// Club 7 Events
			model.NewEvent(43, 7, 6, 10),
			model.NewEvent(44, 7, 11, 15),
			model.NewEvent(45, 7, 16, 20),
			model.NewEvent(46, 7, 21, 25),
			model.NewEvent(47, 7, 26, 30),
			model.NewEvent(48, 7, 31, 35),
			model.NewEvent(49, 7, 36, 40),

			// Club 8 Events
			model.NewEvent(50, 8, 7, 11),
			model.NewEvent(51, 8, 12, 16),
			model.NewEvent(52, 8, 17, 21),
			model.NewEvent(53, 8, 22, 26),
			model.NewEvent(54, 8, 27, 31),
			model.NewEvent(55, 8, 32, 36),
			model.NewEvent(56, 8, 37, 41),

			// Club 9 Events
			model.NewEvent(57, 9, 8, 12),
			model.NewEvent(58, 9, 13, 17),
			model.NewEvent(59, 9, 18, 22),
			model.NewEvent(60, 9, 23, 27),
			model.NewEvent(61, 9, 28, 32),
			model.NewEvent(62, 9, 33, 37),
			model.NewEvent(63, 9, 38, 42),

			// Club 10 Events
			model.NewEvent(64, 10, 9, 13),
			model.NewEvent(65, 10, 14, 18),
			model.NewEvent(66, 10, 19, 23),
			model.NewEvent(67, 10, 24, 28),
			model.NewEvent(68, 10, 29, 33),
			model.NewEvent(69, 10, 34, 38),
			model.NewEvent(70, 10, 39, 43),

			// Club 11-15 Events
			model.NewEvent(71, 11, 10, 14),
			model.NewEvent(72, 11, 15, 19),
			model.NewEvent(73, 11, 20, 24),
			model.NewEvent(74, 11, 25, 29),
			model.NewEvent(75, 11, 30, 34),
			model.NewEvent(76, 11, 35, 39),
			model.NewEvent(77, 11, 40, 44),

			model.NewEvent(78, 12, 11, 15),
			model.NewEvent(79, 12, 16, 20),
			model.NewEvent(80, 12, 21, 25),
			model.NewEvent(81, 12, 26, 30),
			model.NewEvent(82, 12, 31, 35),
			model.NewEvent(83, 12, 36, 40),
			model.NewEvent(84, 12, 41, 45),

			model.NewEvent(85, 13, 12, 16),
			model.NewEvent(86, 13, 17, 21),
			model.NewEvent(87, 13, 22, 26),
			model.NewEvent(88, 13, 27, 31),
			model.NewEvent(89, 13, 32, 36),
			model.NewEvent(90, 13, 37, 41),
			model.NewEvent(91, 13, 42, 46),

			model.NewEvent(92, 14, 13, 17),
			model.NewEvent(93, 14, 18, 22),
			model.NewEvent(94, 14, 23, 27),
			model.NewEvent(95, 14, 28, 32),
			model.NewEvent(96, 14, 33, 37),
			model.NewEvent(97, 14, 38, 42),
			model.NewEvent(98, 14, 43, 47),

			model.NewEvent(99, 15, 14, 18),
			model.NewEvent(100, 15, 19, 23),
		},
	}
}
