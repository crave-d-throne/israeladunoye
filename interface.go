package gopro

type SocialMedia interface {
	Feed() []string
	Fame() int
}

type Facebook struct{
	URL string
	Name string
	Friends int
}
type Twitter struct{
	URL string
	Username string
	Followers int
}
type Linkedin struct{
	URL string
	Name string
	Connections int
}

func (f *Facebook) Feed() []string{
	return []string{
		"Facebook is awesome",
		"i hate it sha",
		"whatever this is, it shall pass"
	}
}
func (f *Facebook) Fame() int{
	return f.Friends
}

func (t *Twitter) Feed(){
	return []string{
		"i like twitter sha",
		"it is just awesome",
		"whatever"
	}
}
func (t *Twitter) Fame(){
	return t.Followers
}

func (l *Linkedin) Feed(){
	return []string{
		"whatever",
		"park i sun",
		"101 dalmatians"
	}
}

func (l *Linkedin) Fame(){
	return l.Connections
}