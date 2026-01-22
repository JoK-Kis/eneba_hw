import { useState, useEffect } from 'react';
import { Search, X } from 'lucide-react';

interface SearchTableProps {
    onSearchChange: (query: string) => void;
}

export default function SearchTable({ onSearchChange }: SearchTableProps) {
    const [searchValue, setSearchValue] = useState('');

    useEffect(() => {
        const timer = setTimeout(() => {
            onSearchChange(searchValue);
        }, 300);

        return () => clearTimeout(timer);
    }, [searchValue, onSearchChange]);

    const handleClear = () => {
        setSearchValue('');
    };

    return (
        <div className="flex items-center border border-white h-full w-100 justify-between px-2">
            <Search className="w-6 h-6 text-white" />
            <input
                type="text"
                value={searchValue}
                onChange={(e) => setSearchValue(e.target.value)}
                placeholder="Search products..."
                className="flex-1 bg-transparent text-white placeholder-gray-400 outline-none px-2"
            />
            {searchValue && (
                <X 
                    className="w-6 h-6 text-white cursor-pointer hover:text-gray-300" 
                    onClick={handleClear}
                />
            )}
        </div>
    );
}