import Lithuania from '../../assets/lithuania.svg';

export default function Language_Currency_Selector() {
    return (
        <div className="flex items-center gap-2">
            <img src={Lithuania} alt="Lithuania" className="h-4 w-4" />
            <p className="text-white text-sm font-light">English EU</p>
            <p className="text-white text-sm font-light">|</p>
            <p className="text-white text-sm font-light">EUR</p>
        </div>
    );
}