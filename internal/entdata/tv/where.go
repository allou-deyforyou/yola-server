// Code generated by entc, DO NOT EDIT.

package tv

import (
	"yola/internal/entdata/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Logo applies equality check predicate on the "logo" field. It's identical to LogoEQ.
func Logo(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// Video applies equality check predicate on the "video" field. It's identical to VideoEQ.
func Video(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVideo), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v bool) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCountry), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// Language applies equality check predicate on the "language" field. It's identical to LanguageEQ.
func Language(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLanguage), v))
	})
}

// LogoEQ applies the EQ predicate on the "logo" field.
func LogoEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLogo), v))
	})
}

// LogoNEQ applies the NEQ predicate on the "logo" field.
func LogoNEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLogo), v))
	})
}

// LogoIn applies the In predicate on the "logo" field.
func LogoIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLogo), v...))
	})
}

// LogoNotIn applies the NotIn predicate on the "logo" field.
func LogoNotIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLogo), v...))
	})
}

// LogoGT applies the GT predicate on the "logo" field.
func LogoGT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLogo), v))
	})
}

// LogoGTE applies the GTE predicate on the "logo" field.
func LogoGTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLogo), v))
	})
}

// LogoLT applies the LT predicate on the "logo" field.
func LogoLT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLogo), v))
	})
}

// LogoLTE applies the LTE predicate on the "logo" field.
func LogoLTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLogo), v))
	})
}

// LogoContains applies the Contains predicate on the "logo" field.
func LogoContains(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLogo), v))
	})
}

// LogoHasPrefix applies the HasPrefix predicate on the "logo" field.
func LogoHasPrefix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLogo), v))
	})
}

// LogoHasSuffix applies the HasSuffix predicate on the "logo" field.
func LogoHasSuffix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLogo), v))
	})
}

// LogoEqualFold applies the EqualFold predicate on the "logo" field.
func LogoEqualFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLogo), v))
	})
}

// LogoContainsFold applies the ContainsFold predicate on the "logo" field.
func LogoContainsFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLogo), v))
	})
}

// VideoEQ applies the EQ predicate on the "video" field.
func VideoEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldVideo), v))
	})
}

// VideoNEQ applies the NEQ predicate on the "video" field.
func VideoNEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldVideo), v))
	})
}

// VideoIn applies the In predicate on the "video" field.
func VideoIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldVideo), v...))
	})
}

// VideoNotIn applies the NotIn predicate on the "video" field.
func VideoNotIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldVideo), v...))
	})
}

// VideoGT applies the GT predicate on the "video" field.
func VideoGT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldVideo), v))
	})
}

// VideoGTE applies the GTE predicate on the "video" field.
func VideoGTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldVideo), v))
	})
}

// VideoLT applies the LT predicate on the "video" field.
func VideoLT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldVideo), v))
	})
}

// VideoLTE applies the LTE predicate on the "video" field.
func VideoLTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldVideo), v))
	})
}

// VideoContains applies the Contains predicate on the "video" field.
func VideoContains(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldVideo), v))
	})
}

// VideoHasPrefix applies the HasPrefix predicate on the "video" field.
func VideoHasPrefix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldVideo), v))
	})
}

// VideoHasSuffix applies the HasSuffix predicate on the "video" field.
func VideoHasSuffix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldVideo), v))
	})
}

// VideoEqualFold applies the EqualFold predicate on the "video" field.
func VideoEqualFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldVideo), v))
	})
}

// VideoContainsFold applies the ContainsFold predicate on the "video" field.
func VideoContainsFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldVideo), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v bool) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v bool) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCountry), v))
	})
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCountry), v))
	})
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCountry), v...))
	})
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCountry), v...))
	})
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCountry), v))
	})
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCountry), v))
	})
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCountry), v))
	})
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCountry), v))
	})
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCountry), v))
	})
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCountry), v))
	})
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCountry), v))
	})
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCountry), v))
	})
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCountry), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// LanguageEQ applies the EQ predicate on the "language" field.
func LanguageEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLanguage), v))
	})
}

// LanguageNEQ applies the NEQ predicate on the "language" field.
func LanguageNEQ(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLanguage), v))
	})
}

// LanguageIn applies the In predicate on the "language" field.
func LanguageIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldLanguage), v...))
	})
}

// LanguageNotIn applies the NotIn predicate on the "language" field.
func LanguageNotIn(vs ...string) predicate.Tv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Tv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldLanguage), v...))
	})
}

// LanguageGT applies the GT predicate on the "language" field.
func LanguageGT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLanguage), v))
	})
}

// LanguageGTE applies the GTE predicate on the "language" field.
func LanguageGTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLanguage), v))
	})
}

// LanguageLT applies the LT predicate on the "language" field.
func LanguageLT(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLanguage), v))
	})
}

// LanguageLTE applies the LTE predicate on the "language" field.
func LanguageLTE(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLanguage), v))
	})
}

// LanguageContains applies the Contains predicate on the "language" field.
func LanguageContains(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLanguage), v))
	})
}

// LanguageHasPrefix applies the HasPrefix predicate on the "language" field.
func LanguageHasPrefix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLanguage), v))
	})
}

// LanguageHasSuffix applies the HasSuffix predicate on the "language" field.
func LanguageHasSuffix(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLanguage), v))
	})
}

// LanguageEqualFold applies the EqualFold predicate on the "language" field.
func LanguageEqualFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLanguage), v))
	})
}

// LanguageContainsFold applies the ContainsFold predicate on the "language" field.
func LanguageContainsFold(v string) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLanguage), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tv) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tv) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tv) predicate.Tv {
	return predicate.Tv(func(s *sql.Selector) {
		p(s.Not())
	})
}
