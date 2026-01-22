import mario from '../assets/product/mario.png';
import splitNintendo from '../assets/product/splitNintendo.jpg';
import fifa from '../assets/product/fifa.jpg';
import fifa2 from '../assets/product/fifa2.jpg';
import fifaXbox from '../assets/product/fifaXbox.jpg';
import fifaXbox2 from '../assets/product/fifaXbox2.jpg';
import rd from '../assets/product/rd.jpg';
import rdbundle from '../assets/product/rdbundle.jpg';
import rdSpecial from '../assets/product/rdSpecial.jpg';
import rdXbox from '../assets/product/rdXbox.jpg';
import split from '../assets/product/split.jpg';

import nintendoIcon from '../assets/platform/nintendo.svg';
import steamIcon from '../assets/platform/steam.svg';
import xboxIcon from '../assets/platform/xbox.svg';
import psnIcon from '../assets/platform/psn.svg';
import eaIcon from '../assets/platform/EA.svg';
import originIcon from '../assets/platform/origin.svg';
import rockstarIcon from '../assets/platform/rockstar_social_club.svg';

export const productImages: Record<string, string> = {
  'mario.png': mario,
  'splitNintendo.jpg': splitNintendo,
  'fifa.jpg': fifa,
  'fifa2.jpg': fifa2,
  'fifaXbox.jpg': fifaXbox,
  'fifaXbox2.jpg': fifaXbox2,
  'rd.jpg': rd,
  'rdbundle.jpg': rdbundle,
  'rdSpecial.jpg': rdSpecial,
  'rdXbox.jpg': rdXbox,
  'split.jpg': split,
};

export const platformIcons: Record<string, string> = {
  'nintendo.svg': nintendoIcon,
  'steam.svg': steamIcon,
  'xbox.svg': xboxIcon,
  'psn.svg': psnIcon,
  'EA.svg': eaIcon,
  'origin.svg': originIcon,
  'rockstar_social_club.svg': rockstarIcon,
};

export const getProductImage = (imagePath: string): string => {
  const filename = imagePath.split('/').pop() || imagePath;
  return productImages[filename] || mario;
};

export const getPlatformIcon = (iconPath: string): string => {
  const filename = iconPath.split('/').pop() || iconPath;
  return platformIcons[filename] || nintendoIcon;
};
