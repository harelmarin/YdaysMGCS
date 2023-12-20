document.body.onload=function() {
    nbr=5;
    p=0;
    container = document.getElementById("container");
    gauche = document.getElementById("gauche")
    droite = document.getElementById("droite")
    container.style.width=(800*nbr)+"px";
    for(i=1;i<=nbr;i++) {
        div=document.createElement("div");
        div.className="photo";
        div.style.backgroundImage="url('')"
        container.appenChild(div);
    }
}