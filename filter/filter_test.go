package filter

import (
	"testing"
)

func TestRunFilters(t *testing.T) {

	t.Run("no filters", func(t *testing.T) {

		var ret = RunFilters(4)

		if ret != 5 {
			t.Errorf("got %d; wanted %d", ret, 5)
		}

	})

	// t.Run("--genres=Comedy", func(t *testing.T) {
	// 	flags := model.ProgramFlags{GenresFlag: "Comedy"}
	// 	genericTest(t, flags, 7, 75)
	// })

}

// func genericTest(t *testing.T, flags model.ProgramFlags, expectedMatches int, expectedHighestLineNumber int) {

// 	file, err := os.Open("../title.basics.truncated.tsv")
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	matchingFileRows, highestLineNumber := RunFilters(scanner, flags)

// 	if len(matchingFileRows) != expectedMatches || highestLineNumber != expectedHighestLineNumber {
// 		t.Errorf("got (%d, %d); wanted (%d, %d)", len(matchingFileRows), highestLineNumber, expectedMatches, expectedHighestLineNumber)
// 	}

// }
