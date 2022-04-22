

//insert one
func Insert{{$.Name}}(session *xorm.Session,model {{$.Name}}) (int64, error) {
	return session.InsertOne(&model)
}

//insert array
func Insert{{$.Name}}s(session *xorm.Session, models []*{{$.Name}}) (int64, error) {
	return session.Insert(&models)
}

//delete
func Delete{{$.Name}}(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (int64, error) {
	for _, where := range wheres {
		session = where(session)
	}
	return session.Delete(&{{$.Name}}{})
}

//update
func Update{{$.Name}}(session *xorm.Session, updates map[string]interface{}, wheres ...func(*xorm.Session) *xorm.Session) (int64, error) {
	session = session.Table(&{{$.Name}}{})
	for _, where := range wheres {
		session = where(session)
	}
	return session.Update(updates)
}

//find
func Find{{$.Name}}(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (result {{$.Name}}, ok bool, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	ok, err = session.Get(&result)
	if err != nil {
		return result,false, err
	}
	return result,ok, nil
}

//finds
func Find{{$.Name}}s(session *xorm.Session, wheres ...func(*xorm.Session) *xorm.Session) (result []*{{$.Name}}, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	if err = session.Find(&result); err != nil {
		return nil, err
	}
	return result, nil
}

//page
func Page{{$.Name}}(session *xorm.Session, page, size int, wheres ...func(*xorm.Session) *xorm.Session) (total int64, result []*{{$.Name}}, err error) {
	for _, where := range wheres {
		session = where(session)
	}
	total, err = session.Limit(size, (page-1)*size).FindAndCount(&result)
	if err != nil {
		return total, nil, err
	}
	return total, result, nil
}

