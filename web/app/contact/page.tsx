import React from 'react';
import { NavigationBarComponent } from '@/components/nav-bar';
import Link from 'next/link';

export default function Contact() {
  return (
    <main className="text-center min-h-screen bg-warm-100">
      <NavigationBarComponent />

      <section className="pt-20 bg-cool-gray-50">
        <h2 className="text-4xl font-bold text-brown-dark-100 mb-10">Contact Us</h2>
        <div className="flex flex-col items-center justify-center space-y-6">
          <form className="contact-form space-y-6 mx-auto max-w-lg shadow-lg p-10 rounded-lg bg-gray">
            <div className="form-field">
              <label htmlFor="firstName" className="block text-gray-700 text-lg mb-2">First Name</label>
              <input
                type="text"
                id="firstName"
                name="firstName"
                className="mt-1 text-gray-700 block w-full rounded-md border-2 border-gray-400 shadow-sm focus:ring-2 focus:ring-pinky-600 focus:border-pinky-600 sm:text-sm p-3"
              />
            </div>
            <div className="form-field">
              <label htmlFor="lastName" className="block text-gray-700 text-lg mb-2">Last Name</label>
              <input
                type="text"
                id="lastName"
                name="lastName"
                className="mt-1 text-gray-700 block w-full rounded-md border-2 border-gray-400 shadow-sm focus:ring-2 focus:ring-pinky-600 focus:border-pinky-600 sm:text-sm p-3"
              />
            </div>
            <div className="form-field">
              <label htmlFor="email" className="block text-gray-700 text-lg mb-2">Email</label>
              <input
                type="email"
                id="email"
                name="email"
                className="mt-1 text-gray-700 block w-full rounded-md border-2 border-gray-400 shadow-sm focus:ring-2 focus:ring-pinky-600 focus:border-pinky-600 sm:text-sm p-3"
              />
            </div>
            <div className="form-field">
              <label htmlFor="message" className="block text-gray-700 text-lg mb-2">Message</label>
              <textarea
                id="message"
                name="message"
                className="mt-1 text-gray-700 block w-full rounded-md border-2 border-gray-400 shadow-sm focus:ring-2 focus:ring-pinky-600 focus:border-pinky-600 sm:text-sm p-3"
                placeholder="Your Message"
              ></textarea>
            </div>
            <button
              type="submit"
              className="w-full inline-flex justify-center py-3 px-6 border border-transparent shadow text-lg font-medium rounded-md text-white bg-pinky-600 hover:bg-pinky-700 focus:outline-none focus:ring-2 focus:ring-pinky-500 transition duration-300 ease-in-out transform hover:-translate-y-1 hover:scale-110"
            >
              Send
            </button>
          </form>
        </div>
      </section>


      <section className="py-8">
        <h2 className="text-3xl font-bold text-brown-dark-100">Follow Us</h2>
        <div className="flex justify-center gap-4 mt-4">
          <Link href="https://www.facebook.com">
            <img src="./icons8-facebook-50.png" alt="" className="" />
          </Link>
          <Link href="https://www.instagram.com/">
            <img src="./icons8-instagram-48.png" alt="" className="" />
          </Link>
        </div>
      </section>
    </main>
  );
};
