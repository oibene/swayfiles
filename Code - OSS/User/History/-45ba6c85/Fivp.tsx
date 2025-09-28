export default function ModalAccount(){
    return(
        <div className="absolute bg-dark-gray w-80 h-50">
            <div className="font-noto font-bold text-light-gray text-sm">

                <div className="m-4">
                    <p>CONTA</p>

                    <div className="flex mt-2">
                        <div className="h-15 w-15 rounded-full bg-light-gray"></div>
                        <div className="ml-4 content-center">
                            <p>Nome Conta</p>
                            <p>email@email.com</p>
                        </div>
                    </div>
                </div>
               
                <hr className="w-full h-0.5 rounded-sm"/>

                <button className="m-4">
                    <p>ALTERAR CONTA</p>
                </button>

                <button>
                    <p>LOGOUT</p>
                </button>
            </div>

        </div>
    )
}