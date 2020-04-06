window.onload = () => {
   const encryptButton = document.getElementById("encryptButton");
   encryptButton.addEventListener("click", encryptAction);

   const decryptButton = document.getElementById("decryptButton"); 
   decryptButton.addEventListener("click", decryptAction);

   const solveButton = document.getElementById("solveButton"); 
   solveButton.addEventListener("click", solveAction);
}

// Handle encrypt button
async function encryptAction() {
   const payload = {
      text: document.getElementById("plaintext-text").value,
      shift: document.getElementById("plaintext-shift").value
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
         document.getElementById("ciphertext-text").value = (await res.json()).response;
      default:
         // Failure
   }
};

// Handle decrypt button
async function decryptAction() { 
   const payload = {
      text: document.getElementById("ciphertext-text").value,
      shift: document.getElementById("ciphertext-shift").value
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
         document.getElementById("plaintext-text").value = (await res.json()).response;
      default:
         // Failure
   }
}

// Handle solve button
async function solveAction() { 
   const payload = {
      text: document.getElementById("ciphertext-text").value,
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
         document.getElementById("plaintext-text").value = (await res.json()).response;
      default:
         // Failure
   }
}