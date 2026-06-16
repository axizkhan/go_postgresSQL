package utils

func GetPagination (page int,limit int)(int32,int32){
	if page<1{
		page=1
	}

	if limit<1{
		limit=1
	}

	if limit>100{
		limit=100
	}

	offeset := (page-1)*limit

	return int32(limit),int32(offeset)
}