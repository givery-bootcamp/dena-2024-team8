export function Signin() {
    return (
        <div className="container-xxl flex justify-center items-center">
            <div className="flex flex-col items-center p-4">
                <input type="text" placeholder="Username" className="mb-4 p-2 border border-gray-300 rounded" />
                <input type="password" placeholder="Password" className="mb-4 p-2 border border-gray-300 rounded" />
                <button type="submit" className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded">
                    Sign In
                </button>
            </div>
        </div>
    );
}
