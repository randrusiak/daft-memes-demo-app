import { API_URL } from "./consts";

const checkURLProtocol = (url) => {
  if (url.includes("https://")) {
    return url;
  } else {
    return url.replace(".", API_URL);
  }
};

export const formatMeme = (meme) => {
  return {
    ...meme,
    imagePath: checkURLProtocol(meme.imagePath),
  };
};

export const mapMemes = (memes) => {
  return memes.map((meme) => formatMeme(meme));
};
