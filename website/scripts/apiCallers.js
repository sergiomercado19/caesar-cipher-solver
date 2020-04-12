import {plainboxText, plainboxShift, cipherboxText, cipherboxShift} from "./elements.js";
import * as verify from "./formVerification.js";

// Handle encrypt button
export async function encryptAction() {
   if (!verify.encryptForm()) return;

   const payload = {
      text: plainboxText.value,
      shift: plainboxShift.value
   }
   const options = {
      method: 'POST',
      headers: {
         'Content-Type': 'application/json'
      },
      body: JSON.stringify(payload)
   }
   const res = await fetch("/encrypt", options);
   switch (res.status) {
      case 200:
         cipherboxText.value = (await res.json()).response;
         cipherboxText.classList.add("output-flash");
      default:
         // Failure
   }
};

// Handle decrypt button
export async function decryptAction() {
   if (!verify.decryptForm()) return;

   const payload = {
      text: cipherboxText.value,
      shift: cipherboxShift.value
   }
   const options = {
      method: 'POST',
      headers: {
         'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload)
   }
   const res = await fetch("/decrypt", options);
   switch (res.status) {
      case 200:
         plainboxText.value = (await res.json()).response;
         plainboxText.classList.add("output-flash");
      default:
         // Failure
   }
}

// Handle solve button
export async function solveAction() {
   if (!verify.solveForm()) return;

   const payload = {
      text: cipherboxText.value,
   }
   const options = {
      method: 'POST',
      headers: {
         'Content-Type': 'application/json',
      },
      body: JSON.stringify(payload)
   }
   const res = await fetch("/solve", options);
   switch (res.status) {
      case 200:
         let ret = await res.json()
         plainboxText.value = ret.response;
         plainboxText.classList.add("output-flash");
         plainboxShift.value = ret.shift;
         plainboxShift.classList.add("output-flash");
      default:
         // Failure
   }
}