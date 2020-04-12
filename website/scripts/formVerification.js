import {plainboxText, plainboxShift, cipherboxText, cipherboxShift} from "./elements.js";

export function encryptForm() {
   const filledPlaintext = plainboxText.value.length != 0;
   const filledShift = plainboxShift.value.length != 0;

   // Remove quotes to prevent JSON parsing issues
   plainboxText.value = plainboxText.value.replace(/['"`]/g, '');

   // Check relevant fields
   if (!filledPlaintext) plainboxText.classList.add("error-flash");
   if (!filledShift) plainboxShift.classList.add("error-flash");

   // Clear flash from non-relevant fields
   cipherboxText.classList.remove("error-flash");
   cipherboxShift.classList.remove("error-flash");
   
   return filledPlaintext && filledShift;
}

export function decryptForm() {
   const filledCiphertext = cipherboxText.value.length != 0;
   const filledShift = cipherboxShift.value.length != 0;

   // Remove quotes to prevent JSON parsing issues
   cipherboxText.value = cipherboxText.value.replace(/['"`]/g, '');

   // Check relevant fields
   if (!filledCiphertext) cipherboxText.classList.add("error-flash");
   if (!filledShift) cipherboxShift.classList.add("error-flash");

   // Clear flash from non-relevant fields
   plainboxText.classList.remove("error-flash");
   plainboxShift.classList.remove("error-flash");

   return filledCiphertext && filledShift;
}

export function solveForm() {
   const filledCiphertext = cipherboxText.value.length != 0;

   // Remove quotes to prevent JSON parsing issues
   cipherboxText.value = cipherboxText.value.replace(/['"`]/g, '');

   // Check relevant fields
   if (!filledCiphertext) cipherboxText.classList.add("error-flash");

   // Clear flash from non-relevant fields
   plainboxText.classList.remove("error-flash");
   plainboxShift.classList.remove("error-flash");
   cipherboxShift.classList.remove("error-flash");

   return filledCiphertext;
}
