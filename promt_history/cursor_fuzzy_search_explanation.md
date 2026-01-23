# Fuzzy search explanation
_Exported on 1/23/2026 at 18:05:08 GMT+2 from Cursor (2.4.21)_

---

**User**

what is fuzzy search

---

**Cursor**

Fuzzy search finds approximate matches even when the query doesn't exactly match the target text.

## How it works

Instead of requiring exact matches, fuzzy search uses similarity metrics to rank results by how close they are to the query. It helps with:
- Typos and misspellings
- Partial matches
- Variations in formatting
- Similar-sounding words

## Common algorithms

1. **Levenshtein distance** — counts character insertions, deletions, and substitutions needed to transform one string into another
2. **Fuzzy matching** — uses similarity scores (0–1) to rank results
3. **N-gram matching** — compares overlapping character sequences

## Example

If you search for "gogle" (typo), fuzzy search might return:
- "Google" (high similarity)
- "Goggle" (high similarity)
- "Gaggle" (lower similarity)

## Common use cases

- Search bars (autocomplete, product search)
- Database queries (finding records with typos)
- Text editors (find-and-replace)
- Data deduplication
- Spell checkers

## In your codebase

If you're working on the Eneba project, fuzzy search could help with product search, allowing users to find products even with typos or partial names.

Want help implementing fuzzy search in your project?

---

**User**

how should i go about implimenting fuzzy search         <p className="text-xs mb-4 font-light">Results found: {products.length}</p>

this has to get the results of search if no search all of available games 

what to do

---

**Cursor**

Reviewing the current code to understand the structure, then implementing fuzzy search.


Checking the search component and backend routes:


Implementing fuzzy search: backend endpoint and frontend integration. Creating the initial todo list:


---

**User**

explain steps ill write code myself

---

**Cursor**

Steps to implement fuzzy search:

## Backend Implementation

### Step 1: Create Search Handler (`backend/internal/handlers/search.go`)
- Accept a `q` query parameter from the request
- If `q` is empty or missing, return all products (same as current `GetProducts`)
- If `q` exists, use PostgreSQL fuzzy matching:
  - Use `ILIKE` for case-insensitive pattern matching
  - Use `%query%` pattern to match anywhere in the product name
  - Example: `WHERE name ILIKE '%query%'`
- Preload the Platform relationship (like in `GetProducts`)
- Return JSON array of matching products

### Step 2: Add Search Route (`backend/cmd/api/main.go`)
- Add a new route: `router.GET("/api/products/search", handlers.SearchProducts)`
- Or modify existing route to accept optional query param: `router.GET("/api/products", handlers.GetProducts)` and handle search in the same handler

## Frontend Implementation

### Step 3: Update SearchTable Component (`frontend/src/components/search/SearchTable.tsx`)
- Add state for search input value (use `useState`)
- Replace static div with an `<input>` element
- Add `onChange` handler to update state as user types
- Optionally add debouncing (wait ~300ms after typing stops before searching)
- Pass search query to parent (App.tsx) via props or context

### Step 4: Update App.tsx (`frontend/src/App.tsx`)
- Add state for search query
- Create a function to fetch products:
  - If search query exists: `fetch('http://localhost:8080/api/products/search?q=' + encodeURIComponent(query))`
  - If no search query: `fetch('http://localhost:8080/api/products')` (current behavior)
- Update `useEffect` to depend on search query
- Pass search query setter to SearchTable component
- The `products.length` will automatically show correct count

## Optional Enhancements

- Debouncing: Wait 300-500ms after user stops typing before making API call
- Loading state: Show loading indicator during search
- Empty state: Show message when no results found
- Search in multiple fields: Search product name, platform name, etc.
- Better fuzzy matching: Use PostgreSQL `pg_trgm` extension for typo tolerance

## Database Query Example (GORM)
```go
// Simple fuzzy search
database.DB.Where("name ILIKE ?", "%"+query+"%").Preload("Platform").Find(&products)

// Or search multiple fields
database.DB.Where("name ILIKE ? OR region ILIKE ?", "%"+query+"%", "%"+query+"%").Preload("Platform").Find(&products)
```

The key is: if no search query → return all products, if search query exists → filter products using fuzzy matching.

