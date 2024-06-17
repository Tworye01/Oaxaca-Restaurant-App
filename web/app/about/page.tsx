import { ImageComponent } from "@/components/image";
import { NavigationBarComponent } from "@/components/nav-bar";

export default function About() {
  return (
    <main className="min-h-screen text-brown-dark-100 bg-warm-100">
      <NavigationBarComponent />

      <div className="pt-40 container mx-auto flex flex-col lg:flex-row items-center py-8 px-4 lg:px-8">
        <div className="flex-1">
          <h1 className="text-3xl font-bold mb-4">About Us</h1>
          <p className="mb-4">
            Welcome to Oaxaca, where the heart of Mexico&apos;s vibrant street food culture thrives in the UK. Born in 2007, our journey from a single London eatery to a beloved chain of 25 is fueled by a passion for unique Mexican flavors and exceptional service. At Oaxaca, we blend the authenticity of Mexican street food with innovative culinary practices, offering a dining experience that&apos;s both familiar and fresh.
          </p>
          <p>
            As we&apos;ve grown, so too has our ambition. Challenged by an evolving market, we are embracing technology to enhance our service while keeping the soul of Oaxaca alive. Our mission? To expand without losing the personal touch that makes us special.
          </p>
          <p className="mt-4">
            Step into Oaxaca for a taste of Mexico&apos;s best, served with a side of innovation. Join us in this exciting chapter where tradition meets tech, and every meal is a celebration.
          </p>
        </div>
        <div className="flex-1 mt-8 lg:mt-0 lg:ml-8">
          <ImageComponent
            src="/vivaPic.jpg"
            alt="Vibrant Mexican culture with a cactus wearing a sombrero and playing guitar"
            width={500}
            height={300}
            objectFit="contain"
          />
        </div>
      </div>
    </main>
  );
}
