export type ServiceCategory = "female" | "male" | "other";

export interface IService {
    id: string;
    barberId: string;
    internalName: string;
    displayName: string;
    category: string;
    description: string;
    durationMinutes: number;
    priceCents: number;
    active: boolean;
    deleted: boolean;
    sortOrder: number;
    createdAt: string;
    updatedAt: string;
}

export const dummyServices: IService[] = [
    {
        id: "11111111-1111-1111-1111-111111111111",
        barberId: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
        internalName: "women_short_haircut",
        displayName: "Frauen Haarschnitt kurz",
        description: "Waschen, Schneiden, Föhnen",
        category: "female",
        durationMinutes: 45,
        priceCents: 3500,
        active: true,
        deleted: false,
        sortOrder: 1,
        createdAt: "2025-01-01T10:00:00Z",
        updatedAt: "2025-01-01T10:00:00Z",
    },
    {
        id: "55555555-5555-5555-5555-555555555555",
        barberId: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
        internalName: "women_long_haircut",
        displayName: "Frauen Haarschnitt lang",
        description: "Waschen, Schneiden, Föhnen",
        category: "female",
        durationMinutes: 60,
        priceCents: 4500,
        active: true,
        deleted: false,
        sortOrder: 2,
        createdAt: "2025-01-01T10:00:00Z",
        updatedAt: "2025-01-01T10:00:00Z",
    },
    {
        id: "22222222-2222-2222-2222-222222222222",
        barberId: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
        internalName: "men_short_haircut",
        displayName: "Herren Haarschnitt kurz",
        description: "Maschine & Schere",
        category: "male",
        durationMinutes: 30,
        priceCents: 2000,
        active: true,
        deleted: false,
        sortOrder: 3,
        createdAt: "2025-01-01T10:00:00Z",
        updatedAt: "2025-01-01T10:00:00Z",
    },
    {
        id: "66666666-6666-6666-6666-666666666666",
        barberId: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
        internalName: "men_long_haircut",
        displayName: "Herren Haarschnitt lang",
        description: "Schere & Styling",
        category: "male",
        durationMinutes: 45,
        priceCents: 3000,
        active: true,
        deleted: false,
        sortOrder: 4,
        createdAt: "2025-01-01T10:00:00Z",
        updatedAt: "2025-01-01T10:00:00Z",
    },
    {
        id: "44444444-4444-4444-4444-444444444444",
        barberId: "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa",
        internalName: "coloring",
        displayName: "Haare Färben",
        description: "Komplettfärbung",
        category: "other",
        durationMinutes: 30,
        priceCents: 6500,
        active: true,
        deleted: false,
        sortOrder: 5,
        createdAt: "2025-01-01T10:00:00Z",
        updatedAt: "2025-01-01T10:00:00Z",
    },
];