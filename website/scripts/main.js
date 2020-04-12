import * as apiCallers from "./apiCallers.js";
import {encryptButton, decryptButton, solveButton} from "./elements.js";
import {plainboxText, plainboxShift, cipherboxText, cipherboxShift} from "./elements.js";

window.onload = () => {
   encryptButton.addEventListener('click', apiCallers.encryptAction);
   decryptButton.addEventListener('click', apiCallers.decryptAction);
   solveButton.addEventListener('click', apiCallers.solveAction);

   plainboxText.addEventListener('input', clearFlash);
   plainboxShift.addEventListener('input', clearFlash);
   cipherboxText.addEventListener('input', clearFlash);
   cipherboxShift.addEventListener('input', clearFlash);
}

function clearFlash(e) {
   e.target.classList.remove('output-flash');
   e.target.classList.remove('error-flash');
}