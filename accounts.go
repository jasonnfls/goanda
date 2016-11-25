package goanda

func AddAccount(a int, token string) error {
	//validate length / error
	Accounts[a] = token
	if len(Accounts) == 1 {
		CurrentAccount = a
	}
	return nil
}

func DeleteAccount(a int) error {

	//check whether aount is current
	if a == CurrentAccount {
		return a2
	}

	//validate existance
	_, ok := Accounts[a]
	if ok {
		delete(Accounts, a)
	} else {
		//return non-existant acount error
		return a1
	}

	//return no problems
	return nil
}

func SetCurrentAccount(a int) error {
	_, ok := Accounts[a]
	if ok {
		CurrentAccount = a
	} else {
		return a1
		//return error for no listed aount
	}
	return nil

}

func ReturnToken(a int) (string, error) {
	token, ok := Accounts[a]
	if ok {
		return token, nil
	} else {
		return "", a3
	}

}
