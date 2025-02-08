function likePost() {
    let count = document.getElementById("likeCount");
    count.innerText = parseInt(count.innerText) + 1; // Increment like count
}

function dislikePost() {
    let count = document.getElementById("dislikeCount");
    count.innerText = parseInt(count.innerText) + 1; // Increment dislike count
}

let userAction = null; // Track user action (null, 'like', 'dislike')

function toggleLike() {
    let likeCount = document.getElementById("likeCount");
    let dislikeCount = document.getElementById("dislikeCount");

    let likeBtn = document.querySelector(".like");
    let dislikeBtn = document.querySelector(".dislike");

    if (userAction === "like") {
        // Undo like
        likeCount.innerText = parseInt(likeCount.innerText) - 1;
        userAction = null;
        likeBtn.classList.remove("active");
    } else {
        // Like the post
        likeCount.innerText = parseInt(likeCount.innerText) + 1;
        likeBtn.classList.add("active");

        // If previously disliked, remove dislike
        if (userAction === "dislike") {
            dislikeCount.innerText = parseInt(dislikeCount.innerText) - 1;
            dislikeBtn.classList.remove("active");
        }

        userAction = "like";
    }
}

function toggleDislike() {
    let likeCount = document.getElementById("likeCount");
    let dislikeCount = document.getElementById("dislikeCount");

    let likeBtn = document.querySelector(".like");
    let dislikeBtn = document.querySelector(".dislike");

    if (userAction === "dislike") {
        // Undo dislike
        dislikeCount.innerText = parseInt(dislikeCount.innerText) - 1;
        userAction = null;
        dislikeBtn.classList.remove("active");
    } else {
        // Dislike the post
        dislikeCount.innerText = parseInt(dislikeCount.innerText) + 1;
        dislikeBtn.classList.add("active");

        // If previously liked, remove like
        if (userAction === "like") {
            likeCount.innerText = parseInt(likeCount.innerText) - 1;
            likeBtn.classList.remove("active");
        }

        userAction = "dislike";
    }
}
