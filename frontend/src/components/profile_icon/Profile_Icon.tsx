import Enebian from '../../assets/enebian-62d31a191efb7.svg';

export default function Profile_Icon() {
    return (
        <div className="border border-white sm:h-8 sm:w-8 h-6 w-6 rounded-full overflow-hidden">
            <img src={Enebian} alt="Enebian" className="h-full w-full object-cover" />
        </div>
    );
}