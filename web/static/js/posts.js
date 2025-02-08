function likePost() {
    let count = document.getElementById("likeCount");
    count.innerText = parseInt(count.innerText) + 1; // Increment like count
}

function dislikePost() {
    let count = document.getElementById("dislikeCount");
    count.innerText = parseInt(count.innerText) + 1; // Increment dislike count
}
