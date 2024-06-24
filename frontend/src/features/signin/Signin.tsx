import React, { useState } from 'react';
import { useAppDispatch, useAppSelector } from '../../shared/hooks';
import { APIService } from '../../shared/services';

export function Signin() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [errorMessage, setErrorMessage] = useState('');
    const dispatch = useAppDispatch();

    const handleSignIn = () => {
        // ログイン処理を実行するコードをここに追加します
        console.log('Username:', username);
        console.log('Password:', password);
        
        // Reduxの呼び出しをしてログイン情報を保存
        dispatch(APIService.signin({ username, password }))
            .then((response) => {
                console.log("response", response);
                // responseからデータを取得し、ログイン成功の処理を行う
                if (response.payload && response.payload) {
                    // homeに遷移
                    window.location.href = "/";
                } else {
                    // データが存在しない場合の処理をここに記述
                    setErrorMessage("ユーザー名またはパスワードが違います。再度お試しください。");
                }
            })
            .catch((error) => {
                // エラー発生時の処理をここに記述
                setErrorMessage("サインインに失敗しました。再度お試しください。"+ error.message);
            });
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
                <p className="mt-4 text-red-500">
                    { errorMessage }
                </p>
            </div>
        </div>
    );
}
