import { useState, useEffect } from 'react';

type Props = {
  selectedFilters: string[];
  setSelectedFilters: (filters: string[]) => void;
  caloriesRange: number[];
  setCaloriesRange: (range: number[]) => void;
}

export const FilterComponent = ({ selectedFilters, setSelectedFilters, caloriesRange, setCaloriesRange }: Props): JSX.Element => {
  const [visible, setVisible] = useState<boolean>(false);

  const toggleFilter = (filter: string) => {
    if (selectedFilters.includes(filter)) {
      setSelectedFilters(selectedFilters.filter(item => item !== filter));
    } else {
      setSelectedFilters([...selectedFilters, filter]);
    }
  };

  const handleRangeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const value = parseInt(event.target.value);
    setCaloriesRange([caloriesRange[0], value]);
  };

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.key === 'Escape') {
        setVisible(false);
      }
    };

    if (visible) {
      document.body.style.overflow = 'hidden';
      document.addEventListener('keydown', handleKeyDown);
    } else {
      document.body.style.overflow = '';
      document.removeEventListener('keydown', handleKeyDown);
    }

    return () => {
      document.body.style.overflow = '';
      document.removeEventListener('keydown', handleKeyDown);
    };
  }, [visible]);

  return (
    <>
      <div className="mt-6 flex justify-start ml-8 ">
        <button onClick={() => setVisible(true)} className="bg-pinky-600 shadow-lg text-black px-4 py-2 mb-4">Filter Menu</button>
      </div>

      {visible && (
        <div className="fixed inset-0 z-50 overflow-hidden backdrop-brightness-50 flex justify-center items-center">
          <div className="bg-white text-brown-dark-100 overflow-hidden shadow-xl max-w-md w-full p-6">

            <h2 className="text-lg text-brown-dark-100 font-semibold mb-4">Filter Menu</h2>

            <div className="space-y-4">
              <div className="flex flex-wrap items-center justify-between">
                <label className="text-sm text-brown-dark-100 font-medium">Allergens:</label>
                <div className="flex flex-wrap gap-2">
                  <button onClick={() => toggleFilter('gluten')} className={`px-3 py-1  text-sm font-medium ${selectedFilters.includes('gluten') ? 'bg-pinky-500 text-gray-800  hover:text-gray-800' : 'bg-gray-200 text-gray-800 hover:bg-pinky-600'}`}>Gluten</button>
                  <button onClick={() => toggleFilter('milk')} className={`px-3 py-1  text-sm font-medium ${selectedFilters.includes('milk') ? 'bg-pinky-500 text-gray-800  hover:text-gray-800' : 'bg-gray-200 text-gray-800 hover:bg-pinky-600'}`}>Milk</button>
                  <button onClick={() => toggleFilter('wheat')} className={`px-3 py-1  text-sm font-medium ${selectedFilters.includes('wheat') ? 'bg-pinky-500 text-gray-800  hover:text-gray-800' : 'bg-gray-200 text-gray-800 hover:bg-pinky-600'}`}>Wheat</button>
                  <button onClick={() => toggleFilter('nut')} className={`px-3 py-1  text-sm font-medium ${selectedFilters.includes('nut') ? 'bg-pinky-500 text-gray-800  hover:text-gray-800' : 'bg-gray-200 text-gray-800 hover:bg-pinky-600'}`}>Nut</button>
                  <button onClick={() => toggleFilter('egg')} className={`px-3 py-1  text-sm font-medium ${selectedFilters.includes('egg') ? 'bg-pinky-500 text-gray-800  hover:text-gray-800' : 'bg-gray-200 text-gray-800 hover:bg-pinky-600'}`}>Egg</button>
                  <button onClick={() => toggleFilter('celery')} className={`px-3 py-1  text-sm font-medium ${selectedFilters.includes('celery') ? 'bg-pinky-500 text-gray-800  hover:text-gray-800' : 'bg-gray-200 text-gray-800 hover:bg-pinky-600'}`}>Celery</button>
                  <button onClick={() => toggleFilter('lupin')} className={`px-3 py-1  text-sm font-medium ${selectedFilters.includes('lupin') ? 'bg-pinky-500 text-gray-800  hover:text-gray-800' : 'bg-gray-200 text-gray-800 hover:bg-pinky-600'}`}>Lupin</button>
                  <button onClick={() => toggleFilter('barley')} className={`px-3 py-1  text-sm font-medium ${selectedFilters.includes('barley') ? 'bg-pinky-500 text-gray-800  hover:text-gray-800' : 'bg-gray-200 text-gray-800 hover:bg-pinky-600'}`}>Barley</button>
                </div>
              </div>

              <div className="flex flex-wrap items-center gap-2">
                <label className="text-sm text-brown-dark-100 font-medium">Vegetarian:</label>
                <button onClick={() => toggleFilter('vegetarian')} className={`px-3 py-1  text-sm font-medium ${selectedFilters.includes('vegetarian') ? 'bg-gradient-to-r from-green-500 to-green-600 text-white' : 'bg-gradient-to-r from-green-100 to-green-200 text-gray-800 hover:bg-gradient-to-r hover:from-green-600 hover:to-green-700 hover:text-white'}`}>Vegetarian</button>
              </div>


              <div>
                <label className="text-sm text-brown-dark-100 font-medium">Calories Range:</label>
                <input
                  type="range"
                  min="0"
                  max="1500"
                  step="100"
                  defaultValue="1500"
                  onChange={handleRangeChange}
                  className="w-full h-2  bg-pinky-600 rounded-lg appearance-none cursor-pointer"
                />
                <span>{caloriesRange[0]} - {caloriesRange[1]} calories</span>
              </div>

            </div>
          </div>
        </div>
      )}

    </>
  );
};
