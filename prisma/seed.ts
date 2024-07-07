import { PrismaClient } from '@prisma/client';

// initialize Prisma Client
const prisma = new PrismaClient();

async function main() {
  // create two dummy articles
  const user1 = await prisma.user.upsert({
    where: { user_name: 'salman' },
    update: {},
    create: {
        user_name: 'salman',
        user_email: 'salman@test.com',
        user_passwordHash:
        "secret",
    },
  });

  const user2 = await prisma.user.upsert({
    where: { user_name: "aissam" },
    update: {},
    create: {
        user_name: 'aissam',
        user_email: 'aissam@test.com',
        user_passwordHash:
        "password",
    },
  });

  console.log({ user1, user2 });
}

// execute the main function
main()
  .catch((e) => {
    console.error(e);
    process.exit(1);
  })
  .finally(async () => {
    // close Prisma Client at the end
    await prisma.$disconnect();
  });
