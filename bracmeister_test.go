package bracmeister

import "testing"

func TestSimple(t *testing.T) {
	if "F" != Calc(168, 88, 56, true).Cup {
		t.Errorf("Decided %s, As fuck", Calc(168, 88, 56, true).Cup)
	}
	if "F" != Calc(168, 88, 56, false).Cup {
		t.Errorf("Decided %s, As fuck", Calc(168, 88, 56, false).Cup)
	}

	if "F" != Calc(162, 90, 58, true).Cup {
		t.Errorf("Decided %s, As fuck", Calc(162, 90, 58, true).Cup)
	}
	if "F" != Calc(162, 90, 58, false).Cup {
		t.Errorf("Decided %s, As fuck", Calc(162, 90, 58, false).Cup)
	}

	if "F" != Calc(162, 88, 58, true).Cup {
		t.Errorf("Decided %s, As fuck", Calc(162, 88, 58, true).Cup)
	}
	if "E" != Calc(162, 88, 58, false).Cup {
		t.Errorf("Decided %s, As fuck", Calc(162, 88, 58, false).Cup)
	}

	if "E" != Calc(161, 86, 58, true).Cup {
		t.Errorf("Decided %s, As fuck", Calc(161, 86, 58, true).Cup)
	}
	if "E" != Calc(161, 86, 58, false).Cup {
		t.Errorf("Decided %s, As fuck", Calc(161, 86, 58, false).Cup)
	}

	if "D" != Calc(160, 84, 57, true).Cup {
		t.Errorf("Decided %s, As fuck", Calc(160, 84, 57, true).Cup)
	}
	if "D" != Calc(160, 84, 57, false).Cup {
		t.Errorf("Decided %s, As fuck", Calc(160, 84, 57, false).Cup)
	}

	if "F" != Calc(170, 92, 62, true).Cup {
		t.Errorf("Decided %s, As fuck", Calc(170, 92, 62, true).Cup)
	}
	if "E" != Calc(170, 92, 62, false).Cup {
		t.Errorf("Decided %s, As fuck", Calc(170, 92, 62, false).Cup)
	}

	if "G" != Calc(164, 93, 59, true).Cup {
		t.Errorf("Decided %s, As fuck", Calc(164, 93, 59, true).Cup)
	}
	if "G" != Calc(164, 93, 59, false).Cup {
		t.Errorf("Decided %s, As fuck", Calc(164, 93, 59, false).Cup)
	}
}
