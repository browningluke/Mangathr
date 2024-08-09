// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/browningluke/mangathr/v2/ent/manga"
	"github.com/browningluke/mangathr/v2/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	mangaFields := schema.Manga{}.Fields()
	_ = mangaFields
	// mangaDescFilteredGroups is the schema descriptor for FilteredGroups field.
	mangaDescFilteredGroups := mangaFields[5].Descriptor()
	// manga.DefaultFilteredGroups holds the default value on creation for the FilteredGroups field.
	manga.DefaultFilteredGroups = mangaDescFilteredGroups.Default.([]string)
	// mangaDescExcludedGroups is the schema descriptor for ExcludedGroups field.
	mangaDescExcludedGroups := mangaFields[6].Descriptor()
	// manga.DefaultExcludedGroups holds the default value on creation for the ExcludedGroups field.
	manga.DefaultExcludedGroups = mangaDescExcludedGroups.Default.([]string)
}
