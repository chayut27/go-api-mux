package models

type Constant struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}

func InsertSuccess() Constant {
	return Constant{
		Result:  true,
		Message: "Insert Success",
	}
}

func InsertError() Constant {
	return Constant{
		Result:  false,
		Message: "Insert Error",
	}
}

func InvalidSyntax() Constant {
	return Constant{
		Result:  false,
		Message: "Invalid Syntax",
	}
}

func Get_data_error() Constant {
	return Constant{
		Result:  false,
		Message: "Get Data Error",
	}
}

func UpdateSuccess() Constant {
	return Constant{
		Result:  true,
		Message: "Update Success",
	}
}

func UpdateError() Constant {
	return Constant{
		Result:  false,
		Message: "Update Error",
	}
}

func DeleteSuccess() Constant {
	return Constant{
		Result:  true,
		Message: "Delete Success",
	}
}

func DeleteError() Constant {
	return Constant{
		Result:  false,
		Message: "Delete Error",
	}
}
