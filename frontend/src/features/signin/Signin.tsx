import React, { useState } from 'react';
import { useAppDispatch, useAppSelector } from '../../shared/hooks';
import { APIService } from '../../shared/services';

export function Signin() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const dispatch = useAppDispatch();

    const handleSignIn = () => {
        // ログイン処理を実行するコードをここに追加します
        console.log('Username:', username);
        console.log('Password:', password);
        
        // Reduxの呼び出しをしてログイン情報を保存
        dispatch(APIService.signin({username, password}));
    };

    return (
        <div className="container-xxl flex justify-center items-center">
            <div className="flex flex-col items-center p-4">
                <input
                    type="text"
                    placeholder="Username"
                    className="mb-4 p-2 border border-gray-300 rounded"
                    value={username}
                    onChange={(e) => setUsername(e.target.value)}
                />
                <input
                    type="password"
                    placeholder="Password"
                    className="mb-4 p-2 border border-gray-300 rounded"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <button
                    type="submit"
                    className="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
                    onClick={handleSignIn}
                >
                    Sign In
                </button>
            </div>
        </div>
    );
}
