import { useState, useEffect } from "react";
import Header from "./components/layout/Header";
import ProductCard from "./components/product/ProductCard";
import { getProductImage, getPlatformIcon } from "./utils/assetMap";

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
  const [searchQuery, setSearchQuery] = useState("");

  useEffect(() => {
    setError(null);

    const API_URL = import.meta.env.VITE_API_URL || "http://localhost:8080";
    const url = searchQuery
      ? `${API_URL}/list?search=${encodeURIComponent(searchQuery)}`
      : `${API_URL}/list`;

    fetch(url, {
      signal: AbortSignal.timeout(80000),
    })
      .then((res) => {
        if (!res.ok) {
          throw new Error("Failed to fetch products");
        }
        return res.json();
      })
      .then((data) => {
        setProducts(data);
        setInitialLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setInitialLoading(false);
        console.error("Error fetching products:", err);
      });
  }, [searchQuery]);

  if (initialLoading) {
    return (
      <div className="min-h-screen bg-background text-white flex items-center justify-center">
        <div className="text-center">
          <p className="text-xl mb-2">Loading products...</p>
          <p className="text-sm text-text-secondary">Backend may take up to 1 minute to start since it's hosted on a free tier service</p>
        </div>
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
    <div className="flex flex-col min-h-screen bg-background text-white px-4 items-center">
      <main className="inline-flex flex-col items-center justify-center">
        <Header onSearchChange={setSearchQuery} />

        <p className="flex sm:justify-start justify-center text-xs font-light mb-4 w-full">
          Results found: {products.length}
        </p>

        <div className="grid xl:grid-cols-4 lg:grid-cols-3 sm:grid-cols-2 grid-cols-1 gap-9">
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
