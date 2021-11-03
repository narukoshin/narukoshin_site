"use strict"

// Setting the links of my social accounts
let socials = {
    "instagram": "https://www.instagram.com/naru.koshin",
    "github": "https://www.github.com/narukoshin",
    "twitter": "https://www.twitter.com/narukoshin"
}

// Redirecting user to the my socials
let moveTo = (name) => {
    if (socials[name]) return window.open(socials[name])
}