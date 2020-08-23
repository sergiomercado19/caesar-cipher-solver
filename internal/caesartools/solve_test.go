package caesartools

import "testing"

func TestSolve(t *testing.T) {

	// Initialize dictionary map
	var dictionary map[string]int = make(map[string]int)
	ParseWords(dictionary)

	t.Run("Solve sentence whose words are found in the dictionary", func(t *testing.T) {
		// var plaintext string = "When one door of happiness closes, another opens; but often we look so long at the closed door that we do not see the one which has been opened for us."
		var cyphertext string = "Epmv wvm lwwz wn pixxqvmaa ktwama, ivwbpmz wxmva; jcb wnbmv em twws aw twvo ib bpm ktwaml lwwz bpib em lw vwb amm bpm wvm epqkp pia jmmv wxmvml nwz ca."
		var shift int = 8

		output := Solve(cyphertext, dictionary)
		if output != shift {
			t.Errorf("Shift doesn't match expected output. Got %v, expected %v", output, shift)
		}
	})

	t.Run("Solve sentence containing words not found in the dictionary", func(t *testing.T) {
		// var plaintext string = "My name is Sergio Mercado Ruiz and I am 21 years old."
		var cyphertext string = "Dp erdv zj Jvixzf Dvitruf Ilzq reu Z rd 21 pvrij fcu."
		var shift int = 17

		output := Solve(cyphertext, dictionary)
		if output != shift {
			t.Errorf("Shift doesn't match expected output. Got %v, expected %v", output, shift)
		}
	})
}
