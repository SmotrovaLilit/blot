import {defineStore} from 'pinia';
import {User} from "@/models/user";
import {v4 as uuidv4} from 'uuid';

export const useUserStore = defineStore('user', {
    state: () => ({
        user: null as User | null,
    }),
    actions: {
        loadUser() {
            const savedUser = localStorage.getItem('user');
            if (savedUser) {
                this.user = JSON.parse(savedUser);
                console.log('load user from local storage', savedUser);
                return;
            }
            console.log('load user from local storage: no user found');
        },
        createUser: function (name: string) {
            let id = uuidv4();
            // TODO: remove this switch
            switch (name) {
                case 'Lilit':
                    id = 'a88bb98e-0e1f-4c82-a209-9d6b4986d013'
                    break;
                case 'Dima':
                    id = 'b88bb98e-0e1f-4c82-a209-9d6b4986d014'
                    break;
                case 'Anya':
                    id = 'c88bb98e-0e1f-4c82-a209-9d6b4986d015'
                    break;
                case 'Armen':
                    id = 'd88bb98e-0e1f-4c82-a209-9d6b4986d016'
                    break;
            }
            this.user = new User(id, name);
            localStorage.setItem('user', JSON.stringify(this.user)); // TODO move to separate function
        },
    },
    getters: {
        userName(): string {
            return this.user?.name || '';
        },
        userId(): string {
            return this.user?.id || '';
        },
    }
});