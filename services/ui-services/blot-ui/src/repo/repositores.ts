import {GrpcGameSetRepository, type GameSetRepository} from "@/repo/grpcGameSetRepository";

const gameSetRemoteRepository = new GrpcGameSetRepository() as GameSetRepository;
export default gameSetRemoteRepository;