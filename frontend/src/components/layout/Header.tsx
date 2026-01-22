import SearchTable from '../search/SearchTable';
import Language_Currency_Selector from '../language_currency_selector/Language_Currency_Selector';
import Profile_Icon from '../profile_icon/Profile_Icon';
import { ShoppingCart, Heart } from 'lucide-react';
import Logo from '../../assets/logoFull.svg';

interface HeaderProps {
    onSearchChange: (query: string) => void;
}

export default function Header({ onSearchChange }: HeaderProps) {
    return (
        <header className="flex justify-between h-20 w-full">
            <div className="flex items-center py-4 w-full justify-between lg:gap-4 gap-2">
                <div className="flex h-full items-center lg:gap-4 gap-2 w-fit">
                    <a href="/" className="h-full flex items-center shrink-0">
                        <img src={Logo} alt="Eneba Logo" className="h-12 w-auto cursor-pointer" />
                    </a>
                    <SearchTable onSearchChange={onSearchChange} />
                    <Language_Currency_Selector />
                </div>

                <div className="flex items-center lg:gap-4 gap-1 hidden sm:flex">
                    <Heart className="w-6 h-6" />
                    <ShoppingCart className="w-6 h-6" />
                    <Profile_Icon />
                </div>
            </div>

        </header>
    );
}