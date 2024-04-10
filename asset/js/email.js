function sendEmail() {
    const nom = document.getElementById('nom').value;
    const prenom = document.getElementById('prenom').value;
    const email = document.getElementById('email').value;
    const message = document.getElementById('message').value;

    Email.send({
        SecureToken: "C973D7AD-F097-4B95-91F4-40ABC5567812",
        To: 'them@website.com',
        From: email, // Utilisation de l'e-mail de l'utilisateur comme expéditeur
        Subject: "Nouveau message de " + nom + " " + prenom,
        Body: "Message de : " + nom + " " + prenom + "\n\n" + message
    }).then(
        message => alert("E-mail envoyé avec succès!")
    );
}
