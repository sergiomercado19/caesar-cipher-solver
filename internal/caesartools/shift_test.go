package caesartools

import "testing"

func TestShiftBy5(t *testing.T) {
	var plaintext string = "A Caesar cipher is one of the simplest and most widely known encryption techniques."
	var cyphertext string = "F Hfjxfw hnumjw nx tsj tk ymj xnruqjxy fsi rtxy bnijqd pstbs jshwduynts yjhmsnvzjx."
	var shift int = 5

	t.Run("Shift forward by 5", func(t *testing.T) {
		output := Shift(plaintext, shift)
		if output != cyphertext {
			t.Errorf("Output doesn't match expected output\n=== Shifted by %d ===\n%v\n=== Expected ===\n%v", shift, output, cyphertext)
		}
	})

	t.Run("Shift backward by 5", func(t *testing.T) {
		output := Shift(cyphertext, -shift)
		if output != plaintext {
			t.Errorf("Output doesn't match expected output\n=== Shifted by %d ===\n%v\n=== Expected ===\n%v", shift, output, plaintext)
		}
	})
}

func TestShiftBy21(t *testing.T) {
	var plaintext string = "It is a type of substitution cipher in which each letter in the plaintext is replaced by a letter some fixed number of positions down the alphabet."
	var cyphertext string = "Do dn v otkz ja npwnodopodji xdkczm di rcdxc zvxc gzoozm di ocz kgvdiozso dn mzkgvxzy wt v gzoozm njhz adszy iphwzm ja kjndodjin yjri ocz vgkcvwzo."
	var shift int = 21

	t.Run("Shift forward by 21", func(t *testing.T) {
		output := Shift(plaintext, shift)
		if output != cyphertext {
			t.Errorf("Output doesn't match expected output\n=== Shifted by %d ===\n%v\n=== Expected ===\n%v", shift, output, cyphertext)
		}
	})

	t.Run("Shift backward by 21", func(t *testing.T) {
		output := Shift(cyphertext, -shift)
		if output != plaintext {
			t.Errorf("Output doesn't match expected output\n=== Shifted by %d ===\n%v\n=== Expected ===\n%v", shift, output, plaintext)
		}
	})
}

func TestShiftByCompleteCycle(t *testing.T) {
	var plaintext string = "The Caesar cipher is named after Julius Caesar, who used it with a shift of three (3) to protect messages of military significance."
	var cycle int = 26

	t.Run("Shift forward by 0 cycles", func(t *testing.T) {
		output := Shift(plaintext, 0)
		if output != plaintext {
			t.Errorf("Output doesn't match expected output\n=== Shifted by %d ===\n%v\n=== Expected ===\n%v", 0, output, plaintext)
		}
	})

	t.Run("Shift forward by 1 cycle", func(t *testing.T) {
		output := Shift(plaintext, cycle)
		if output != plaintext {
			t.Errorf("Output doesn't match expected output\n=== Shifted by %d ===\n%v\n=== Expected ===\n%v", cycle, output, plaintext)
		}
	})

	t.Run("Shift backward by 1 cycle", func(t *testing.T) {
		output := Shift(plaintext, -cycle)
		if output != plaintext {
			t.Errorf("Output doesn't match expected output\n=== Shifted by %d ===\n%v\n=== Expected ===\n%v", cycle, output, plaintext)
		}
	})
}
