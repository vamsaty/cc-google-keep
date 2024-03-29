import {Inter} from "next/font/google";
import "./globals.css";
import {ProtectedApp} from "././page";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Google keep",
  description: "Generated by create next app",
};

export default function RootLayout({ children }) {
    return (
        <html lang="en">
            <body className={inter.className}>
                <ProtectedApp>
                    {children}
                </ProtectedApp>
            </body>
        </html>
    );
}
