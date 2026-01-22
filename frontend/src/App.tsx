import { useState, useEffect } from 'react';
import Header from './components/layout/Header';
import ProductCard from './components/product/ProductCard';
import { getProductImage, getPlatformIcon } from './utils/assetMap';

interface Platform {
  id: number;
  name: string;
  icon: string;
}

interface Product {
  id: number;
  name: string;
  image: string;
  platform_id: number;
  platform: Platform;
  location: string;
  originalPrice: number;
  currentPrice: number;
  discount: number;
  cashbackAmount: number;
  likes: number;
  inStock: boolean;
}

function App() {
  const [products, setProducts] = useState<Product[]>([]);
  const [initialLoading, setInitialLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [searchQuery, setSearchQuery] = useState('');

  useEffect(() => {
    setError(null);
    
    const url = searchQuery 
      ? `http://localhost:8080/api/search?q=${encodeURIComponent(searchQuery)}`
      : 'http://localhost:8080/api/search';

    fetch(url)
      .then(res => {
        if (!res.ok) {
          throw new Error('Failed to fetch products');
        }
        return res.json();
      })
      .then(data => {
        setProducts(data);
        setInitialLoading(false);
      })
      .catch(err => {
        setError(err.message);
        setInitialLoading(false);
        console.error('Error fetching products:', err);
      });
  }, [searchQuery]);

  if (initialLoading) {
    return (
      <div className="min-h-screen bg-background text-white flex items-center justify-center">
        <p className="text-xl">Loading products...</p>
      </div>
    );
  }

  if (error) {
    return (
      <div className="min-h-screen bg-background text-white flex items-center justify-center">
        <p className="text-xl text-red-500">Error: {error}</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-background text-white">
      <Header onSearchChange={setSearchQuery} />
      
      <main className="max-w-5xl mx-auto py-4">
        <p className="text-xs mb-4 font-light">Results found: {products.length}</p>

        <div className="grid grid-cols-4 gap-9">
          {products.map((product) => (
            <ProductCard
              key={product.id}
              image={getProductImage(product.image)}
              name={product.name}
              location={product.location}
              platformIcon={getPlatformIcon(product.platform.icon)}
              platformName={product.platform.name}
              originalPrice={product.originalPrice}
              currentPrice={product.currentPrice}
              discount={product.discount}
              cashbackAmount={product.cashbackAmount}
              likes={product.likes}
            />
          ))}
        </div>
      </main>
    </div>
  );
}

export default App;