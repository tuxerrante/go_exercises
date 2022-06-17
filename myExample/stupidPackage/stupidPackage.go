package stupidPackage
import f "fmt"

func Hello(s *string) (string, error){
	if s == nil || *s == "" {
		return "", f.Errorf("ERROR: no string given")
	}
	return f.Sprintf("Helloooooo %s", *s), nil
}
