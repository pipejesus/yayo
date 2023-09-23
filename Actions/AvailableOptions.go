package Actions

func DbOptions() ActionOptions {
	return ActionOptions{
		ActionOption{
			"mariadb",
			"MariaDB",
		},
		ActionOption{
			"mysql:5.7",
			"MySQL 5.7",
		},
		ActionOption{
			"mysql:8.0",
			"MySQL 8.0",
		},
	}
}

func PhpOptions() ActionOptions {
	return ActionOptions{
		ActionOption{
			"8.2",
			"PHP 8.2",
		},
		ActionOption{
			"8.1",
			"PHP 8.1",
		},
		ActionOption{
			"8.0",
			"PHP 8.0",
		},
		ActionOption{
			"7.4",
			"PHP 7.4",
		},
		ActionOption{
			"7.3",
			"PHP 7.3",
		},
		ActionOption{
			"7.2",
			"PHP 7.2",
		},
	}
}
