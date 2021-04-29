package ministrystats

// Category defines the ministry statistics category.
type Category string

// Categories defines a collection of Category.
type Categories []Category

const endPointPrefix = "minhealth"

func validCategories() Categories {
	return Categories{"pharmacists", "pharmacies", "doctors", "dentists"}
}

func (c Category) isValid() bool {
	for _, vc := range validCategories() {
		if vc == c {
			return true
		}
	}

	return false
}
