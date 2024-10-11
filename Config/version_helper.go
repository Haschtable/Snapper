package config

import "fmt"

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewVersion(major, minor, patch int) Version {
	return Version{major, minor, patch}
}

func (v Version) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func StringToVersion(version string) Version {
	var major, minor, patch int
	fmt.Sscanf(version, "v%d.%d.%d", &major, &minor, &patch)
	return Version{major, minor, patch}
}

func (v Version) Compare(other Version) int {
	if v.Major != other.Major {
		return v.Major - other.Major
	}
	if v.Minor != other.Minor {
		return v.Minor - other.Minor
	}
	if v.Patch != other.Patch {
		return v.Patch - other.Patch
	}
	return 0
}

func (v Version) LessThan(other Version) bool {
	return v.Compare(other) < 0
}

func (v Version) LessThanOrEqual(other Version) bool {
	return v.Compare(other) <= 0
}

func (v Version) Min(other Version) bool {
	return !v.LessThan(other)
}
