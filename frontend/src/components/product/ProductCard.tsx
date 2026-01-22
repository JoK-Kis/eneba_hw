import { Heart, InfoIcon } from "lucide-react";
import Refresh from '../../assets/refresh.svg';

interface ProductCardProps {
    image: string;
    name: string;
    location: string;
    platformIcon: string;
    platformName: string;
    originalPrice: number;
    currentPrice: number;
    discount: number;
    cashbackAmount: number;
    likes: number;
}

export default function ProductCard({
    image,
    name,
    location,
    platformIcon,
    platformName,
    originalPrice,
    currentPrice,
    discount,
    cashbackAmount,
    likes
}: ProductCardProps) {

    return (
        <div className="flex-col h-[480px] w-[225px] border-secondary border overflow-hidden group relative cursor-pointer">
            <img src={image} alt={name} className="w-full h-7/11 object-cover" />

            <div className="flex flex-col transition-transform duration-300 group-hover:-translate-y-24 -mt-[22px]">
                {cashbackAmount > 0 && (
                    <div className="flex -mt-11 mb-3 justify-center items-center w-25 h-8 bg-[#63e3c1] gap-1">
                        <img src={Refresh} alt="Cashback" className="w-4 h-4" />
                        <p className="text-primary text-xs font-semibold">CASHBACK</p>
                    </div>
                )}

                <div className="flex h-[22px] bg-black/60 backdrop-blur-sm justify-center items-center gap-1">
                    <img src={platformIcon} alt={platformName} className="w-4 h-4" />
                    <p className="text-white text-[10px] font-light">{platformName}</p>
                </div>

                <div className="flex flex-col justify-between w-full bg-primary p-3 min-h-[271px]">
                    <div className="flex flex-col gap-1 min-h-[60px]">
                        <p className="text-white text-xs font-light line-clamp-2">{name}</p>
                        <p className="text-location text-xs font-semibold">{location}</p>
                    </div>

                    <div className="flex flex-col min-h-[60px]">
                        <div className="flex items-center gap-1">
                            <p className="text-text-secondary text-xs font-semibold">From</p>
                            {discount > 0 && (
                                <p className="text-text-secondary text-xs font-semibold line-through">€{originalPrice.toFixed(2)}</p>
                            )}
                            {discount > 0 && (
                                <p className="text-cashback text-xs font-semibold">-{discount}%</p>
                            )}
                        </div>
                        <div className="flex items-center gap-1">
                            <p className="text-white text-xl font-bold">€{currentPrice.toFixed(2)}</p>
                            <InfoIcon className="w-5 h-5 text-text-secondary cursor-pointer" />
                        </div>
                        {cashbackAmount > 0 && (
                            <div className="flex items-center gap-1">
                                <p className="text-cashback text-xs font-medium">Cashback:</p>
                                <p className="text-cashback text-xs font-medium">€{cashbackAmount.toFixed(2)}</p>
                            </div>
                        )}
                    </div>

                    <div className="flex items-center gap-1">
                        <Heart className="w-4 h-4 text-text-secondary" />
                        <p className="text-text-secondary text-xs font-medium">{likes}</p>
                    </div>

                    <button className="w-full bg-cta text-primary text-sm font-semibold py-2 cursor-pointer hover:bg-cta/80 transition-all duration-300">Add to Cart</button>
                    <button className="w-full border text-white text-sm font-semibold py-2 cursor-pointer hover:border-cta transition-all duration-300 hover:bg-black/30">Explore options</button>
                </div>
            </div>

        </div>
    );
}