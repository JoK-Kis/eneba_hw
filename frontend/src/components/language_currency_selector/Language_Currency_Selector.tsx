import Lithuania from '../../assets/lithuania.svg';

export default function Language_Currency_Selector() {
    return (
        <div className="lg:flex items-center lg:gap-2 gap-1 w-full hidden">
            <img src={Lithuania} alt="Lithuania" className="h-4 w-4" />
            <p className="text-white lg:text-sm text-xs font-light">English EU</p>
            <p className="text-white lg:text-sm text-xs font-light">|</p>
            <p className="text-white lg:text-sm text-xs font-light">EUR</p>
        </div>
    );
}