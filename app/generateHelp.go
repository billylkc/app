package app

type GroupCommand struct {
	Group    string
	Order    int
	Commands []Command
}

type Command struct {
	Name        string // Name of the command. Daily Sales
	Short       string // Short way to call the command, e.g. app d s
	Long        string // Long way to call the command, e.g. app daily sales
	Advanced    string // Another way call the command, usually with extra param, e.g. app daily sales -d "2021-03-03"
	Description string // Description
}

func GetCommandStruct() []GroupCommand {
	var cc []GroupCommand

	// Daily
	daily := []Command{
		Command{
			Name:        "Sales",
			Short:       "app d s 0 7",
			Long:        "app daily sales 0 7",
			Advanced:    "app daily sales -d `2021-03-03`",
			Description: "Daily Sales for the past n days",
		},
		Command{
			Name:        "Products",
			Short:       "app d p 0 7",
			Long:        "app daily products 0 7",
			Advanced:    "app daily products -d `2021-03-03`",
			Description: "Daily Products for the past n days",
		},
		Command{
			Name:        "Members",
			Short:       "app d m 0 7",
			Long:        "app daily members 0 7",
			Advanced:    "app daily members -d `2021-03-03`",
			Description: "Daily Members for the past n days",
		},
		Command{
			Name:        "Refund",
			Short:       "app d r 0 7",
			Long:        "app daily refund 0 7",
			Advanced:    "app daily refund -d `2021-03-03`",
			Description: "Daily Refund amount for the past n days",
		},
	}

	// Weekly
	weekly := []Command{
		Command{
			Name:        "Sales",
			Short:       "app w s 0 7",
			Long:        "app weekly sales 0 4",
			Advanced:    "app weekly sales -d `2021-03-03`",
			Description: "Weekly Sales for the past n weeks",
		},
		Command{
			Name:        "Products",
			Short:       "app w p 0 4",
			Long:        "app weekly products 0 4",
			Advanced:    "app weekly products -d `2021-03-03`",
			Description: "Weekly Products for the past n weeks",
		},
		Command{
			Name:        "Members",
			Short:       "app w m",
			Long:        "app weekly members 0 4",
			Advanced:    "app weekly members -d `2021-03-03`",
			Description: "Weekly Members for the past n weeks",
		},
		Command{
			Name:        "Refund",
			Short:       "app w r 0 7",
			Long:        "app weekly refund 0 7",
			Advanced:    "app weekly refund -d `2021-03-03`",
			Description: "Weekly Refund amount for the past n weeks",
		},
	}

	// Monthly
	monthly := []Command{
		Command{
			Name:        "Sales",
			Short:       "app m s 0 7",
			Long:        "app monthly sales 0 4",
			Advanced:    "app monthly sales -d `2021-03-01`",
			Description: "Monthly Sales for the past n months",
		},
		Command{
			Name:        "Products",
			Short:       "app m p 0 4",
			Long:        "app monthly products 0 4",
			Advanced:    "app monthly products -d `2021-03-01`",
			Description: "Monthly Products for the past n months",
		},
		Command{
			Name:        "Members",
			Short:       "app m m",
			Long:        "app monthly members 0 4",
			Advanced:    "app monthly members -d `2021-03-01`",
			Description: "Monthly Members for the past n months",
		},
		Command{
			Name:        "Refund",
			Short:       "app m r 0 7",
			Long:        "app monthly refund 0 4",
			Advanced:    "app monthly refund -d `2021-03-01`",
			Description: "Monthly Refund amount for the past n month",
		},
	}

	// User
	user := []Command{
		Command{
			Name:        "New",
			Short:       "app u n 0 6",
			Long:        "app user new 0 4",
			Advanced:    "app use new -d `2021-03-01`",
			Description: "No. of new users per month.",
		},
		Command{
			Name:        "Paid",
			Short:       "app u p 0 6",
			Long:        "app user paid 0 4",
			Advanced:    "app use paid -d `2021-03-01`",
			Description: "No. of paid users per month.",
		},
	}

	// Top
	top := []Command{
		Command{
			Name:        "Members",
			Short:       "app t m",
			Long:        "app top members 10",
			Advanced:    "app top members -m jan 10",
			Description: "All time (Or after certain month) top spending members.",
		},
		Command{
			Name:        "Products",
			Short:       "app t p",
			Long:        "app top products 10",
			Advanced:    "app top products -m jan 10",
			Description: "All time (Or after certain month) top selling products.",
		},
	}

	// Query
	query := []Command{
		Command{
			Name:        "Member",
			Short:       "app q m Jimbelle Yu",
			Long:        "app query member Jimbelle Yu",
			Advanced:    "app top member -m jan Jimbelle Yu",
			Description: "Query individual members buying history.",
		},
		Command{
			Name:        "Product",
			Short:       "app q p 177976",
			Long:        "app query product 177976",
			Advanced:    "app query product 177976",
			Description: "Query individual product order history.",
		},
	}

	// Append all
	cc = []GroupCommand{
		GroupCommand{
			Group:    "[d] Daily",
			Order:    1,
			Commands: daily,
		},
		GroupCommand{
			Group:    "[w] Weekly",
			Order:    2,
			Commands: weekly,
		},
		GroupCommand{
			Group:    "[m] Monthly",
			Order:    3,
			Commands: monthly,
		},
		GroupCommand{
			Group:    "[u] User",
			Order:    4,
			Commands: user,
		},
		GroupCommand{
			Group:    "[t] Top",
			Order:    5,
			Commands: top,
		},
		GroupCommand{
			Group:    "[q] Query",
			Order:    6,
			Commands: query,
		},
	}
	return cc
}
