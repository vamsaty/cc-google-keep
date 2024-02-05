'use client';

import {AuthProvider, doLogin, doRegister} from "./auth/auth";

export default function Home() {
  return (
      <>
          <AuthProvider.Context>
            <main className="flex min-h-screen flex-col items-center justify-between p-24">
                <button onClick={(e)=>doLogin('user', 'user')}>Login</button>
                <button onClick={(e)=>doRegister('user', 'user')}>Register</button>
            </main>
          </AuthProvider.Context>
      </>
  );
}
