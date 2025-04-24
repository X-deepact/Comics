import comicController from './comics';
import genreController from './genres';
import adsController from './ads';
import chaptersController from './chapters';
import commonController from './common';

export const API = {
  comic: comicController,
  genre: genreController,
  ads: adsController,
  chapter: chaptersController,
  service: commonController,
};
