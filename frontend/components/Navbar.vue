<template>
  <header>
    <nav class="bg-white border-gray-200 px-4 lg:px-6 py-2.5 dark:bg-gray-800">
      <div class="flex flex-wrap justify-between items-center mx-auto max-w-screen-xl">
        <nuxt-link to="/" class="flex items-center">
          <img src="https://static.vecteezy.com/system/resources/previews/023/546/330/original/shiba-inu-dog-in-business-suit-and-sunglasses-free-png.png" class="mr-3 h-6 sm:h-9" alt="Bros Logo" />
          <span class="self-center text-xl font-semibold whitespace-nowrap dark:text-white">jESuS brOs</span>
        </nuxt-link>
        <div class="flex items-center lg:order-2">
          <nuxt-link to="/" class="text-gray-800 dark:text-white hover:bg-gray-50 focus:ring-4 focus:ring-gray-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:hover:bg-gray-700 focus:outline-none dark:focus:ring-gray-800">Log in</nuxt-link>
          <nuxt-link to="/newbro" class="text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-300 font-medium rounded-lg text-sm px-4 lg:px-5 py-2 lg:py-2.5 mr-2 dark:bg-primary-600 dark:hover:bg-primary-700 focus:outline-none dark:focus:ring-primary-800">New Bro</nuxt-link>
          <button
            data-collapse-toggle="mobile-menu-2"
            type="button"
            class="inline-flex items-center p-2 ml-1 text-sm text-gray-500 rounded-lg lg:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-600"
            aria-controls="mobile-menu-2"
            aria-expanded="false"
          >
            <!-- Button contents -->
          </button>
        </div>
        <div
          class="hidden justify-between items-center w-full lg:flex lg:w-auto lg:order-1"
          id="mobile-menu-2"
        >
          <ul class="flex flex-col mt-4 font-medium lg:flex-row lg:space-x-8 lg:mt-0">
            <li>
              <nuxt-link
                to="/"
                class="block py-2 pr-4 pl-3 text-white rounded bg-primary-700 lg:bg-transparent lg:text-primary-700 lg:p-0 dark:text-white"
                aria-current="page"
              >Home</nuxt-link>
            </li>
            <li>
              <nuxt-link
                to="/about"
                class="block py-2 pr-4 pl-3 text-gray-700 border-b border-gray-100 hover:bg-gray-50 lg:hover:bg-transparent lg:border-0 lg:hover:text-primary-700 lg:p-0 dark:text-gray-400 lg:dark:hover-text-white dark:hover-bg-gray-700 dark:hover-text-white lg:dark:hover-bg-transparent dark:border-gray-700"
              >about</nuxt-link>
            </li>
            <!-- Add more links for other pages -->
          </ul>
        </div>
      </div>
    </nav>
  </header>
</template>
